package groupchat

import (
	"encoding/json"
	"net/http"

	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/modules/message"
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/utils"

	"github.com/gorilla/websocket"
)

type Handler struct {
	groupChatService  GroupChatService
	messageRepository message.MessageRepository
}

type connection struct {
	UserID  uint64
	Conn    *websocket.Conn
	GroupID uint64
}

type input struct {
	ID         uint64 `json:"id"`
	UserID     uint64 `json:"userId"`
	Name       string `json:"name"`
	Members    []entity.User
	PictureUrl string `json:"pictureUrl"`
}

func NewHandler(groupChatService GroupChatService, messageRepository message.MessageRepository) Handler {
	return Handler{groupChatService, messageRepository}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")

		if origin == "http://localhost:3000" {
			return true
		} else {
			return false
		}
	},
}

var connections = make(map[uint64][]*connection)

func (h *Handler) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Failed to upgrade to WebSocket", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	groupID := utils.GetOneQueryParam(w, r, "groupId", "number").(uint64)
	userID := utils.GetOneQueryParam(w, r, "userId", "number").(uint64)

	existingConnections := connections[groupID]
	for _, existingConn := range existingConnections {
		if existingConn.UserID == userID {
			existingConn.Conn.Close()
			connections[groupID] = removeConnection(groupID, existingConn)
			break
		}
	}

	newConn := &connection{
		UserID:  userID,
		Conn:    conn,
		GroupID: groupID,
	}
	connections[groupID] = append(connections[groupID], newConn)

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			connections[groupID] = removeConnection(groupID, newConn)
			return
		}

		newMessage, err := h.messageRepository.Create(userID, 0, groupID, string(message))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		broadcastMessage(groupID, newMessage)
	}
}

func broadcastMessage(groupID uint64, newMessage *entity.Message) {
	for _, conn := range connections[groupID] {
		err := conn.Conn.WriteJSON(newMessage)
		if err != nil {
			connections[groupID] = removeConnection(groupID, conn)
		}
	}
}

func removeConnection(groupID uint64, connToRemove *connection) []*connection {
	conns := connections[groupID]
	for i, conn := range conns {
		if conn == connToRemove {
			return append(conns[:i], conns[i+1:]...)
		}
	}
	return conns
}

func (h *Handler) CreateGroupChat(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	group, _ := h.groupChatService.CreateGroupChat(input.Name, input.Members)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(group); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) AddMembers(w http.ResponseWriter, r *http.Request) {
	groupId := utils.GetPathParam(w, r, "groupId", "number").(uint64)

	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	group, _ := h.groupChatService.AddMembers(groupId, input.Members)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(group); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetAllGroupChats(w http.ResponseWriter, r *http.Request) {
	userId := utils.GetPathParam(w, r, "userId", "number").(uint64)

	feedback, _ := h.groupChatService.GetAllGroupChats(userId)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(feedback); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetOneGroupChat(w http.ResponseWriter, r *http.Request) {
	id := utils.GetPathParam(w, r, "id", "number").(uint64)

	group, _ := h.groupChatService.GetOneGroupChat(id)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(group); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) UpdateGroupChat(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	feedback, _ := h.groupChatService.UpdateGroupChat(input.ID, input.PictureUrl)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(feedback); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) LeaveGroupChat(w http.ResponseWriter, r *http.Request) {

	userId := utils.GetPathParam(w, r, "userId", "number").(uint64)
	groupId := utils.GetPathParam(w, r, "groupId", "number").(uint64)

	err := h.groupChatService.LeaveGroupChat(userId, groupId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	resp := struct {
		Message string `json:"message"`
	}{
		Message: "request success",
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func (h *Handler) DeleteGroupChat(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.groupChatService.DeleteGroupChat(input.ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	resp := struct {
		Message string `json:"message"`
	}{
		Message: "request success",
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
