package aIconversation

import (
	"net/http"

	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/utils"
)

type Handler struct {
	aIConversationService AIConversationService
}

func NewHandler(aIConversationService AIConversationService) Handler {
	return Handler{aIConversationService}
}

func (h *Handler) CreateConversation(w http.ResponseWriter, r *http.Request) {
	user, err := utils.DecodeAccessToken(r)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	conversation, err := h.aIConversationService.CreateConversation(user.Id)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(w, conversation)
}

func (h *Handler) GetConversation(w http.ResponseWriter, r *http.Request) {
	user, err := utils.DecodeAccessToken(r)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	
	conversation, _ := h.aIConversationService.GetConversation(user.Id)

	utils.SuccessResponse(w, conversation)
}

func (h *Handler) DeleteConversation(w http.ResponseWriter, r *http.Request) {
	var err error

	conversationId := utils.GetPathParam(r, "conversationId", "number", &err).(uint64)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	err = h.aIConversationService.DeleteConversation(conversationId)

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
