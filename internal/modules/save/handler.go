package save

import (
	"encoding/json"
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/utils"
	"net/http"
)

type Handler struct {
	saveContentService SaveContentService
}

type input struct {
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
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	like, err := h.saveContentService.SaveContent(input.UserID, input.ContentID)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(w, like)
}

func (h *Handler) GetAllSaves(w http.ResponseWriter, r *http.Request) {
	var err error

	userId := utils.GetPathParam(r, "userId", "number", &err).(uint64)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	contents, err := h.saveContentService.GetAllSaves(userId)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(w, contents)
}

func (h *Handler) GetOneSave(w http.ResponseWriter, r *http.Request) {
	var err error

	userId := utils.GetQueryParam(r, "userId", "number", &err).(uint64)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	contentId := utils.GetQueryParam(r, "contentId", "number", &err).(uint64)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	content, err := h.saveContentService.GetOneSave(userId, contentId)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(w, content)
}

func (h *Handler) UnsaveContent(w http.ResponseWriter, r *http.Request) {
	var err error

	id := utils.GetPathParam(r, "id", "number", &err).(uint64)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	err = h.saveContentService.UnsaveContent(id)

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
