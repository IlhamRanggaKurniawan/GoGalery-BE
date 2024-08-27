package notification

import (
	"encoding/json"
	"net/http"

	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/utils"
)

type Handler struct {
	notificationService NotificationService
}

type input struct {
	ReceiverID uint64 `json:"receiverId"`
	TriggerID  uint64 `json:"triggerId"`
	Content    string `json:"content"`
}

func NewHandler(notificationService NotificationService) Handler {
	return Handler{notificationService}
}

func (h *Handler) CreateNotification(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	notification, _ := h.notificationService.CreateNotification(input.ReceiverID, input.TriggerID, input.Content)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(notification); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetAllNotifications(w http.ResponseWriter, r *http.Request) {

	userId := utils.GetPathParam(w, r, "userId", "number").(uint64)

	notifications, _ := h.notificationService.GetAllNotifications(userId)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(notifications); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) UpdateNotifications(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	messages, _ := h.notificationService.UpdateNotifications(input.ReceiverID)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(messages); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DeleteNotifications(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.notificationService.DeleteNotifications(input.ReceiverID)

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
