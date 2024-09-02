package directmessage

import (
	"encoding/json"
	"net/http"

	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/modules/message"
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/utils"
	"github.com/gorilla/websocket"
)

type Handler struct {
	directMessageService DirectMessageService
	messageRepository    message.MessageRepository
}

type input struct {
	ID           uint64   `json:"id"`
	UserID       uint64   `json:"userId"`
	Participants []uint64 `json:"Participants"`
}

type connection struct {
	UserID uint64
	Conn   *websocket.Conn
	DmID   uint64
}

func NewHandler(directMessageService DirectMessageService, messageRepository message.MessageRepository) Handler {
	return Handler{directMessageService, messageRepository}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins (use with caution)
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

	dmID := utils.GetOneQueryParam(w, r, "dmId", "number").(uint64)
	userID := utils.GetOneQueryParam(w, r, "userId", "number").(uint64)

	// Periksa apakah koneksi sudah ada untuk userID dan dmID
	existingConnections := connections[dmID]
	for _, existingConn := range existingConnections {
		if existingConn.UserID == userID {
			existingConn.Conn.Close()  // Tutup koneksi lama
			connections[dmID] = removeConnection(dmID, existingConn)
			break
		}
	}

	newConn := &connection{
		UserID: userID,
		Conn:   conn,
		DmID:   dmID,
	}
	connections[dmID] = append(connections[dmID], newConn)

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			connections[dmID] = removeConnection(dmID, newConn)
			return
		}

		newMessage, err := h.messageRepository.Create(userID, dmID, 0, string(message))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		broadcastMessage(dmID, newMessage)
	}
}

func broadcastMessage(dmId uint64, newMessage *entity.Message) {
	for _, conn := range connections[dmId] {
		err := conn.Conn.WriteJSON(newMessage)
		if err != nil {
			connections[dmId] = removeConnection(dmId, conn)
		}
	}
}

func removeConnection(dmId uint64, connToRemove *connection) []*connection {
	conns := connections[dmId]
	for i, conn := range conns {
		if conn == connToRemove {
			return append(conns[:i], conns[i+1:]...)
		}
	}
	return conns
}

func (h *Handler) CreateDirectMessage(w http.ResponseWriter, r *http.Request) {
	var input input
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	feedback, err := h.directMessageService.CreateDirectMessage(input.Participants)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(feedback); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) GetAllDirectMessages(w http.ResponseWriter, r *http.Request) {
	userId := utils.GetPathParam(w, r, "userId", "number").(uint64)
	if userId == 0 {
		http.Error(w, "params is empty", http.StatusBadRequest)
		return
	}

	directMessage, err := h.directMessageService.GetAllDirectMessages(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(directMessage); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) GetOneDirectMessageByParticipants(w http.ResponseWriter, r *http.Request) {
	params := map[string]string{
		"participant1Id": "number",
		"participant2Id": "number",
	}

	results := utils.GetMultipleQueryParams(w, r, params)
	if results == nil {
		http.Error(w, "Missing participant1 and participant2", http.StatusBadRequest)
		return
	}

	participant1Id, ok := results["participant1Id"].(uint64)
	if !ok {
		http.Error(w, "Invalid type for 'participant1Id'", http.StatusInternalServerError)
		return
	}

	participant2Id, ok := results["participant2Id"].(uint64)
	if !ok {
		http.Error(w, "Invalid type for 'participant2Id'", http.StatusInternalServerError)
		return
	}

	participants := []uint64{participant1Id, participant2Id}

	directMessage, err := h.directMessageService.GetOneDirectMessageByParticipants(participants)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(directMessage); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) GetOneDirectMessage(w http.ResponseWriter, r *http.Request) {
	id := utils.GetPathParam(w, r, "id", "number").(uint64)

	directMessage, err := h.directMessageService.GetOneDirectMessage(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(directMessage); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) DeleteDirectMessage(w http.ResponseWriter, r *http.Request) {
	var input input
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.directMessageService.DeleteDirectMessage(input.ID)
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
	}
}
