package comment

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	commentService CommentService
}

type input struct {
	ID        uint64   `json:"id"`
	ContentID uint64   `json:"contentId"`
	UserID    uint64   `json:"userId"`
	Message   string `json:"message"`
}

func NewHandler(commentService CommentService) Handler {
	return Handler{commentService}
}

func (h *Handler) SendComment(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	content, _ := h.commentService.SendComment(input.UserID, input.ContentID, input.Message)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(content); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	content, _ := h.commentService.GetAllComments(input.ContentID)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(content); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	content, _ := h.commentService.updateComment(input.ID, input.Message)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(content); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.commentService.DeleteContent(input.ID)

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
