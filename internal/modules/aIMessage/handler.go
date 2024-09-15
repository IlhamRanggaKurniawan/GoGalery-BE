package aImessage

import (
	"encoding/json"
	"net/http"

	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/fetchAPI"
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/utils"
)

type Handler struct {
	aIMessageService AIMessageService
}

type input struct {
	SenderID       uint64             `json:"senderId"`
	ConversationId uint64             `json:"conversationId"`
	Prompt         []fetchAPI.Message `json:"prompt"`
}

func NewHandler(aIMessageService AIMessageService) Handler {
	return Handler{aIMessageService}
}

func (h *Handler) SendMessage(w http.ResponseWriter, r *http.Request) {
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

	message, err := h.aIMessageService.SendMessage(input.SenderID, conversationId, input.Prompt)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(w, message)
}

// func (h *Handler) UpdateMessage(w http.ResponseWriter, r *http.Request) {
// 	var err error

// 	messageId := utils.GetPathParam(r, "messageId", "number", &err).(uint64)

// 	var input input

// 	err = json.NewDecoder(r.Body).Decode(&input)

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	message, _ := h.aIMessageService.UpdateMessage(messageId, input.Message)

// 	w.Header().Set("Content-Type", "application/json")

// 	if err := json.NewEncoder(w).Encode(message); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }

func (h *Handler) DeleteMessage(w http.ResponseWriter, r *http.Request) {
	var err error

	messageId := utils.GetPathParam(r, "messageId", "number", &err).(uint64)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	err = h.aIMessageService.DeleteMessage(messageId)

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
