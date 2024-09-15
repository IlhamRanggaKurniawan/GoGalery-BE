package like

import (
	"encoding/json"
	"net/http"

	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/utils"
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
	var err error

	contentId := utils.GetPathParam(r, "contentId", "number", &err).(uint64)

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

	like, err := h.likeContentService.LikeContent(input.UserID, contentId)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(w, like)
}

func (h *Handler) GetOneLike(w http.ResponseWriter, r *http.Request) {
	var err error

	contentId := utils.GetPathParam(r, "contentId", "number", &err).(uint64)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	userId := utils.GetQueryParam(r, "userId", "number", &err).(uint64)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	like, err := h.likeContentService.GetOneLike(userId, contentId)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(w, like)
}

func (h *Handler) UnlikeContent(w http.ResponseWriter, r *http.Request) {
	var err error

	likeId := utils.GetPathParam(r, "likeId", "number", &err).(uint64)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	err = h.likeContentService.UnlikeContent(likeId)

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
