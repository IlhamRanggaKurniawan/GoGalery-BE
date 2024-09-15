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
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	groupID := utils.GetQueryParam(r, "groupId", "number", &err).(uint64)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	userID := utils.GetQueryParam(r, "userId", "number", &err).(uint64)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

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
			utils.ErrorResponse(w, err, http.StatusInternalServerError)
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
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	group, err := h.groupChatService.CreateGroupChat(input.Name, input.Members)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(w, group)
}

func (h *Handler) AddMembers(w http.ResponseWriter, r *http.Request) {
	var err error

	groupId := utils.GetPathParam(r, "groupId", "number", &err).(uint64)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	var input input

	err = json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	group, err := h.groupChatService.AddMembers(groupId, input.Members)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(w, group)
}

func (h *Handler) GetAllGroupChats(w http.ResponseWriter, r *http.Request) {
	var err error

	userId := utils.GetPathParam(r, "userId", "number", &err).(uint64)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	groups, err := h.groupChatService.GetAllGroupChats(userId)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(w, groups)
}

func (h *Handler) GetOneGroupChat(w http.ResponseWriter, r *http.Request) {
	var err error

	groupId := utils.GetPathParam(r, "groupId", "number", &err).(uint64)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	group, err := h.groupChatService.GetOneGroupChat(groupId)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(w, group)
}

func (h *Handler) UpdateGroupChat(w http.ResponseWriter, r *http.Request) {
	var err error

	groupId := utils.GetPathParam(r, "groupId", "number", &err).(uint64)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	var input input

	err = json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	group, err := h.groupChatService.UpdateGroupChat(groupId, input.PictureUrl)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(w, group)
}

func (h *Handler) LeaveGroupChat(w http.ResponseWriter, r *http.Request) {
	var err error

	userId := utils.GetPathParam(r, "userId", "number", &err).(uint64)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	groupId := utils.GetPathParam(r, "groupId", "number", &err).(uint64)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	err = h.groupChatService.LeaveGroupChat(userId, groupId)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	resp := struct {
		Message string `json:"message"`
	}{
		Message: "request success",
	}

	utils.SuccessResponse(w, resp)
}
func (h *Handler) DeleteGroupChat(w http.ResponseWriter, r *http.Request) {
	var err error

	groupId := utils.GetPathParam(r, "groupId", "number", &err).(uint64)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	err = h.groupChatService.DeleteGroupChat(groupId)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	resp := struct {
		Message string `json:"message"`
	}{
		Message: "request success",
	}

	utils.SuccessResponse(w, resp)
}
