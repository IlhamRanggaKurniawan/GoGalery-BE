package save

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	saveContentService SaveContentService
}

type input struct {
	ID        uint `json:"id"`
	UserID    uint `json:"userId"`
	ContentID uint `json:"contentId"`
}

func NewHandler(saveContentService SaveContentService) Handler {
	return Handler{saveContentService}
}

func (h *Handler) SaveContent(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	like, _ := h.saveContentService.SaveContent(input.UserID, input.ContentID)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(like); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetAllSaves(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	likes, _ := h.saveContentService.GetAllSaves(input.ContentID)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(likes); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetOneSave(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	likes, _ := h.saveContentService.GetOneSave(input.UserID, input.ContentID)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(likes); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) UnsaveContent(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.saveContentService.UnsaveContent(input.ID)

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
