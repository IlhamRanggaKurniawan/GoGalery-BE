package save

import (
	"encoding/json"
	"net/http"
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/utils"
)

type Handler struct {
	saveContentService SaveContentService
}

type input struct {
	ID        uint64 `json:"id"`
	UserID    uint64 `json:"userId"`
	ContentID uint64 `json:"contentId"`
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
	params := map[string]string{
		"userId":    "number",
		"contentId": "number",
	}

	results := utils.GetMultipleQueryParams(w, r, params)

	userId, ok := results["userId"].(uint64)
	if !ok {
		http.Error(w, "Invalid type for 'userId'", http.StatusBadRequest)
		return
	}

	contentId, ok := results["contentId"].(uint64)
	if !ok {
		http.Error(w, "Invalid type for 'contentId'", http.StatusBadRequest)
		return
	}

	likes, _ := h.saveContentService.GetOneSave(userId, contentId)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(likes); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) UnsaveContent(w http.ResponseWriter, r *http.Request) {
	id := utils.GetPathParam(w, r, "id", "number").(uint64)

	if id == 0 {
		http.Error(w, "params is empty", http.StatusBadRequest)
		return
	}

	err := h.saveContentService.UnsaveContent(id)

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
