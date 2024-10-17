package notification

import (
	"net/http"

	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/utils"
)

type Handler struct {
	notificationService NotificationService
}

func NewHandler(notificationService NotificationService) Handler {
	return Handler{notificationService}
}

func (h *Handler) GetAllNotifications(w http.ResponseWriter, r *http.Request) {
	user, err := utils.DecodeAccessToken(r)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	_, err = h.notificationService.GetAllNotifications(user.Id)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	notifications, err := h.notificationService.UpdateNotifications(user.Id)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(w, notifications)
}

func (h *Handler) DeleteNotifications(w http.ResponseWriter, r *http.Request) {
	user, err := utils.DecodeAccessToken(r)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	err = h.notificationService.DeleteNotifications(user.Id)

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
