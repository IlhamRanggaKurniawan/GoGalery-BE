package like

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	likeContentService LikeContentService
}

type input struct {
	ID        uint64 `json:"id"`
	UserID    uint64 `json:"userId"`
	ContentID uint64 `json:"contentId"`
}

func NewHandler(likeContentService LikeContentService) Handler {
	return Handler{likeContentService}
}

func (h *Handler) LikeContent(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	like, _ := h.likeContentService.LikeContent(input.UserID, input.ContentID)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(like); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetAllLikes(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	likes, _ := h.likeContentService.GetAllLikes(input.ContentID)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(likes); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetOneLike(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	likes, _ := h.likeContentService.GetOneLike(input.UserID, input.ContentID)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(likes); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) UnlikeContent(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.likeContentService.UnlikeContent(input.ID)

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
