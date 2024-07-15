package groupchat

import (
	"encoding/json"
	"net/http"

	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"
)

type Handler struct {
	groupChatService GroupChatService
}

type input struct {
	ID         uint64 `json:"id"`
	UserID     uint64 `json:"userId"`
	Name       string `json:"name"`
	Members    []entity.User
	PictureUrl string `json:"pictureUrl"`
}

func NewHandler(groupChatService GroupChatService) Handler {
	return Handler{groupChatService}
}

func (h *Handler) CreateGroupChat(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	feedback, _ := h.groupChatService.CreateGroupChat(input.Name, input.Members)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(feedback); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetAllGroupChats(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	feedback, _ := h.groupChatService.GetAllGroupChats(input.UserID)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(feedback); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetOneGroupChat(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	feedback, _ := h.groupChatService.GetOneGroupChat(input.ID)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(feedback); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) UpdateGroupChat(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	feedback, _ := h.groupChatService.UpdateGroupChat(input.ID, input.PictureUrl)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(feedback); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DeleteGroupChat(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.groupChatService.DeleteGroupChat(input.ID)

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
