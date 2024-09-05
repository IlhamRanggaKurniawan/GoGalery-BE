package aImessage

import (
	"encoding/json"
	"net/http"

	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/fetchAPI"
)

type Handler struct {
	aIMessageService AIMessageService
}

type input struct {
	ID             uint64          `json:"id"`
	SenderID       uint64          `json:"senderId"`
	ConversationId uint64          `json:"conversationId"`
	Prompt         []fetchAPI.Message `json:"prompt"`
}

func NewHandler(aIMessageService AIMessageService) Handler {
	return Handler{aIMessageService}
}

func (h *Handler) SendMessage(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	message, _ := h.aIMessageService.SendMessage(input.SenderID, input.ConversationId, input.Prompt)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(message); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetAllMessages(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	messages, _ := h.aIMessageService.GetAllMessages(input.ConversationId)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(messages); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// func (h *Handler) UpdateMessage(w http.ResponseWriter, r *http.Request) {
// 	var input input

// 	err := json.NewDecoder(r.Body).Decode(&input)

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	message, _ := h.aIMessageService.UpdateMessage(input.ID, input.Message)

// 	w.Header().Set("Content-Type", "application/json")

// 	if err := json.NewEncoder(w).Encode(message); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }

func (h *Handler) DeleteMessage(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.aIMessageService.DeleteMessage(input.ID)

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
