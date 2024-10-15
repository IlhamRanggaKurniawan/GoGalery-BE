package directmessage

import (
	"encoding/json"
	"net/http"
	"os"

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
	UserId       uint64   `json:"userId"`
	Participants []uint64 `json:"Participants"`
}

type connection struct {
	UserId uint64
	Conn   *websocket.Conn
	DmId   uint64
}

func NewHandler(directMessageService DirectMessageService, messageRepository message.MessageRepository) Handler {
	return Handler{directMessageService, messageRepository}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		if origin == os.Getenv("FRONT_END_ORIGIN") {
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

	dmId := utils.GetPathParam(r, "dmId", "number", &err).(uint64)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	user, err := utils.DecodeAccessToken(r)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	existingConnections := connections[dmId]

	for _, existingConn := range existingConnections {
		if existingConn.UserId == user.Id {
			existingConn.Conn.Close()
			connections[dmId] = removeConnection(dmId, existingConn)
			break
		}
	}

	newConn := &connection{
		UserId: user.Id,
		Conn:   conn,
		DmId:   dmId,
	}

	connections[dmId] = append(connections[dmId], newConn)

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			connections[dmId] = removeConnection(dmId, newConn)
			return
		}

		newMessage, err := h.messageRepository.Create(user.Id, dmId, 0, string(message))
		if err != nil {
			utils.ErrorResponse(w, err, http.StatusInternalServerError)
			return
		}

		broadcastMessage(dmId, newMessage)
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
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	
	user, err := utils.DecodeAccessToken(r)
	
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	input.Participants = append(input.Participants, user.Id)

	directMessage, err := h.directMessageService.CreateDirectMessage(input.Participants)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(w, directMessage)
}

func (h *Handler) GetAllDirectMessages(w http.ResponseWriter, r *http.Request) {
	user, err := utils.DecodeAccessToken(r)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	directMessages, err := h.directMessageService.GetAllDirectMessages(user.Id)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(w, directMessages)
}

func (h *Handler) GetOneDirectMessageByParticipants(w http.ResponseWriter, r *http.Request) {
	user, err := utils.DecodeAccessToken(r)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	userId := utils.GetPathParam(r, "userId", "number", &err).(uint64)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	participants := []uint64{user.Id, userId}

	directMessage, err := h.directMessageService.GetOneDirectMessageByParticipants(participants)

	utils.SuccessResponse(w, directMessage)
}

func (h *Handler) GetOneDirectMessage(w http.ResponseWriter, r *http.Request) {
	var err error

	dmId := utils.GetPathParam(r, "dmId", "number", &err).(uint64)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	directMessage, err := h.directMessageService.GetOneDirectMessage(dmId)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(w, directMessage)
}

func (h *Handler) DeleteDirectMessage(w http.ResponseWriter, r *http.Request) {
	var err error

	dmId := utils.GetPathParam(r, "dmId", "number", &err).(uint64)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	err = h.directMessageService.DeleteDirectMessage(dmId)

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
