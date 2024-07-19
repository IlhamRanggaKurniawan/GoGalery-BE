package content

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	contentService ContentService
}

type input struct {
	ID         uint64 `json:"id"`
	UploaderID uint64 `json:"uploaderId"`
	UserID     uint64 `json:"userId"`
	Caption    string `json:"caption"`
}

func NewHandler(contentService ContentService) Handler {
	return Handler{contentService}
}

func (h *Handler) UploadContent(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	content, _ := h.contentService.UploadContent(input.UploaderID, input.Caption, input.Caption)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(content); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) UpdateContent(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	content, _ := h.contentService.UpdateContent(input.ID, input.Caption)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(content); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetAllContent(w http.ResponseWriter, r *http.Request) {

	contents, _ := h.contentService.GetAllContents()

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(contents); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetOneContent(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	content, _ := h.contentService.GetOneContent(input.ID)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(content); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetAllContentByFollowing(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	content, _ := h.contentService.GetAllContentsByFollowing(input.UserID)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(content); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DeleteContent(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.contentService.DeleteContent(input.ID)

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
