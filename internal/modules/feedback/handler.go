package feedback

import (
	"encoding/json"
	"net/http"

	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/utils"
)

type Handler struct {
	feedbackService FeedbackService
}

type input struct {
	Message string `json:"message"`
}

func NewHandler(feedbackService FeedbackService) Handler {
	return Handler{feedbackService}
}

func (h *Handler) SendFeedback(w http.ResponseWriter, r *http.Request) {
	var err error

	userId := utils.GetPathParam(r, "userId", "number", &err).(uint64)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	var input input

	err = json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	feedback, err := h.feedbackService.SendFeedback(userId, input.Message)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(w, feedback)
}

func (h *Handler) GetAllFeedbacks(w http.ResponseWriter, r *http.Request) {

	feedbacks, err := h.feedbackService.GetAllFeedbacks()

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(w, feedbacks)
}
