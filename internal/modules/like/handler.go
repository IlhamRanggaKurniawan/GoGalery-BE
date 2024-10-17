package like

import (
	"net/http"

	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/utils"
)

type Handler struct {
	likeContentService LikeContentService
}

func NewHandler(likeContentService LikeContentService) Handler {
	return Handler{likeContentService}
}

func (h *Handler) LikeContent(w http.ResponseWriter, r *http.Request) {
	user, err := utils.DecodeAccessToken(r)
	
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	
	contentId := utils.GetPathParam(r, "contentId", "number", &err).(uint64)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	like, err := h.likeContentService.LikeContent(user.Id, contentId)

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

	user, err := utils.DecodeAccessToken(r)

	like, err := h.likeContentService.GetOneLike(user.Id, contentId)

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
