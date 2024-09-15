package message

import (
	"encoding/json"
	"net/http"

	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/utils"
)

type Handler struct {
	messageService MessageService
}

type input struct {
	SenderID        uint64 `json:"senderId"`
	Message         string `json:"message"`
	DirectMessageID uint64 `json:"directMessageId"`
	GroupChatID     uint64 `json:"groupChatId"`
}

func NewHandler(messageService MessageService) Handler {
	return Handler{messageService}
}

func (h *Handler) SendPrivateMessage(w http.ResponseWriter, r *http.Request) {
	var err error

	conversationId := utils.GetPathParam(r, "conversationId", "number", &err).(uint64)

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

	message, err := h.messageService.SendMessage(input.SenderID, conversationId, 0, input.Message)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(w, message)
}

func (h *Handler) SendGroupMessage(w http.ResponseWriter, r *http.Request) {
	var err error

	conversationId := utils.GetPathParam(r, "conversationId", "number", &err).(uint64)

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

	message, err := h.messageService.SendMessage(input.SenderID, 0, conversationId, input.Message)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(w, message)
}

func (h *Handler) UpdateMessage(w http.ResponseWriter, r *http.Request) {
	var err error

	messageId := utils.GetPathParam(r, "messageId", "number", &err).(uint64)

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

	message, err := h.messageService.UpdateMessage(messageId, input.Message)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(w, message)
}

func (h *Handler) DeleteMessage(w http.ResponseWriter, r *http.Request) {
	var err error

	messageId := utils.GetPathParam(r, "messageId", "number", &err).(uint64)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	err = h.messageService.DeleteMessage(messageId)

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
