package aIconversation

import (
	"encoding/json"
	"net/http"

	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/utils"
)

type Handler struct {
	aIConversationService AIConversationService
}

type input struct {
	ID     uint64 `json:"id"`
	UserID uint64 `json:"userId"`
}

func NewHandler(aIConversationService AIConversationService) Handler {
	return Handler{aIConversationService}
}

func (h *Handler) CreateConversation(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	conversation, _ := h.aIConversationService.CreateConversation(input.UserID)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(conversation); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetConversation(w http.ResponseWriter, r *http.Request) {
	userId := utils.GetPathParam(w, r, "userId", "number").(uint64)

	if userId == 0 {
		http.Error(w, "params is empty", http.StatusBadRequest)
		return
	}

	conversation, _ := h.aIConversationService.GetConversation(userId)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(conversation); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DeleteConversation(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.aIConversationService.DeleteConversation(input.ID)

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
