package comment

import (
	"encoding/json"
	"net/http"

	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/utils"
)

type Handler struct {
	commentService CommentService
}

type input struct {
	ContentID uint64 `json:"contentId"`
	UserID    uint64 `json:"userId"`
	Message   string `json:"message"`
}

func NewHandler(commentService CommentService) Handler {
	return Handler{commentService}
}

func (h *Handler) SendComment(w http.ResponseWriter, r *http.Request) {
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

	content, err := h.commentService.SendComment(input.UserID, contentId, input.Message)

	utils.SuccessResponse(w, content)
}

func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	var err error

	contentId := utils.GetPathParam(r, "contentId", "number", &err).(uint64)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	content, err := h.commentService.GetAllComments(contentId)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(w, content)
}

func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	var err error

	commentId := utils.GetPathParam(r, "commentId", "number", &err).(uint64)

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

	content, err := h.commentService.updateComment(commentId, input.Message)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(w, content)
}

func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	var err error

	commentId := utils.GetPathParam(r, "commentId", "number", &err).(uint64)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	err = h.commentService.DeleteContent(commentId)

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
