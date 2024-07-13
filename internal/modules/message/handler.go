package message

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	messageService MessageService
}

type input struct {
	ID              uint   `json:"id"`
	SenderID        uint   `json:"senderId"`
	Text            string `json:"text"`
	DirectMessageID uint   `json:"directMessageId"`
	GroupChatID     uint   `json:"groupChatId"`
}

func NewHandler(messageService MessageService) Handler {
	return Handler{messageService}
}

func (h *Handler) SendMessage(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	message, _ := h.messageService.SendMessage(input.SenderID, input.DirectMessageID, input.GroupChatID, input.Text)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(message); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetAllMessage(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	messages, _ := h.messageService.GetAllMessages(input.DirectMessageID, input.GroupChatID)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(messages); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) UpdateMessage(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	messages, _ := h.messageService.UpdateMessage(input.ID, input.Text)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(messages); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DeleteMessage(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.messageService.DeleteMessage(input.ID)

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
