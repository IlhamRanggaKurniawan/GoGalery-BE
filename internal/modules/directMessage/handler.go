package directmessage

import (
	"encoding/json"
	"net/http"
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/utils"
)

type Handler struct {
	directMessageService DirectMessageService
}

type input struct {
	ID           uint64   `json:"id"`
	UserID       uint64   `json:"userId"`
	Participants []uint64 `json:"Participants"`
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
	userId := utils.GetPathParam(w, r, "userId", "number").(uint64)

	if userId == 0 {
		http.Error(w, "params is empty", http.StatusBadRequest)
		return
	}

	directMessage, _ := h.directMessageService.GetAllDirectMessages(userId)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(directMessage); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetOneDirectMessageByParticipants(w http.ResponseWriter, r *http.Request) {
	params := map[string]string{
		"participant1Id": "number",
		"participant2Id": "number",
	}

	results := utils.GetMultipleQueryParams(w, r, params)

	if results == nil {
		http.Error(w, "Missing participant1 and participant2", http.StatusBadRequest)
		return
	}

	participant1Id, ok := results["participant1Id"].(uint64)
	if !ok {
		http.Error(w, "Invalid type for 'id'", http.StatusInternalServerError)
		return
	}

	participant2Id, ok := results["participant2Id"].(uint64)
	if !ok {
		http.Error(w, "Invalid type for 'name'", http.StatusInternalServerError)
		return
	}

	participants := []uint64{participant1Id, participant2Id}

	directMessage, _ := h.directMessageService.GetOneDirectMessageByParticipants(participants)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(directMessage); err != nil {
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
