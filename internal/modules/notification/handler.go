package notification

import (
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

func (h *Handler) GetAllNotifications(w http.ResponseWriter, r *http.Request) {
	var err error

	receiverId := utils.GetPathParam(r, "receiverId", "number", &err).(uint64)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	notifications, err := h.notificationService.GetAllNotifications(receiverId)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	notifications, err = h.notificationService.UpdateNotifications(receiverId)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(w, notifications)
}

func (h *Handler) DeleteNotifications(w http.ResponseWriter, r *http.Request) {
	var err error

	receiverId := utils.GetPathParam(r, "receiverId", "number", &err).(uint64)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	err = h.notificationService.DeleteNotifications(receiverId)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	resp := struct {
		Message string `json:"message"`
	}{
		Message: "request success",
	}

	utils.SuccessResponse(w, resp)
}
