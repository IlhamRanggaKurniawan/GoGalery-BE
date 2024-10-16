package save

import (
	"net/http"

	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/utils"
)

type Handler struct {
	saveContentService SaveContentService
}

func NewHandler(saveContentService SaveContentService) Handler {
	return Handler{saveContentService}
}

func (h *Handler) SaveContent(w http.ResponseWriter, r *http.Request) {
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

	save, err := h.saveContentService.SaveContent(user.Id, contentId)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(w, save)
}

func (h *Handler) GetAllSaves(w http.ResponseWriter, r *http.Request) {
	user, err := utils.DecodeAccessToken(r)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	contents, err := h.saveContentService.GetAllSaves(user.Id)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(w, contents)
}

func (h *Handler) GetOneSave(w http.ResponseWriter, r *http.Request) {
	user, err := utils.DecodeAccessToken(r)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	contentId := utils.GetQueryParam(r, "contentId", "number", &err).(uint64)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	content, err := h.saveContentService.GetOneSave(user.Id, contentId)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(w, content)
}

func (h *Handler) UnsaveContent(w http.ResponseWriter, r *http.Request) {
	var err error

	saveId := utils.GetPathParam(r, "saveId", "number", &err).(uint64)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	err = h.saveContentService.UnsaveContent(saveId)

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
