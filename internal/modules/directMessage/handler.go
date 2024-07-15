package directmessage

import (
	"encoding/json"
	"net/http"

	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"
)

type Handler struct {
	directMessageService DirectMessageService
}

type input struct {
	ID           uint64 `json:"id"`
	UserID       uint64 `json:"userId"`
	Participants []entity.User
}

func NewHandler(directMessageService DirectMessageService) Handler {
	return Handler{directMessageService}
}

func (h *Handler) CreateDirectMessage(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	feedback, _ := h.directMessageService.CreateDirectMessage(input.Participants)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(feedback); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetAllDirectMessages(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	feedback, _ := h.directMessageService.GetAllDirectMessages(input.UserID)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(feedback); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetOneDirectMessage(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	feedback, _ := h.directMessageService.GetOneDirectMessage(input.ID)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(feedback); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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
		return
	}
}
