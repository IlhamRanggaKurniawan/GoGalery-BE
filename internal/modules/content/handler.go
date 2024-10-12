package content

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/modules/like"
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/modules/save"
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/utils"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Handler struct {
	contentService ContentService
	likeService    like.LikeContentService
	saveService    save.SaveContentService
	S3Client       *s3.Client
	BucketName     string
}

type input struct {
	Id         uint64 `json:"id"`
	UploaderId uint64 `json:"uploaderId"`
	UserId     uint64 `json:"userId"`
	Caption    string `json:"caption"`
	Path       string `json:"path"`
}

type Like struct {
	IsLiked bool   `json:"isLiked"`
	LikeId  uint64 `json:"likeId"`
}

type Save struct {
	IsSaved bool   `json:"isSaved"`
	SaveId  uint64 `json:"saveId"`
}

type ContentResponse struct {
	Content entity.Content `json:"content"`
	Like    Like
	Save    Save
}

type ContentListResponse struct {
	Contents   []ContentResponse `json:"contents"`
	NextCursor *uint64           `json:"nextCursor"`
}

func NewHandler(service ContentService, s3Client *s3.Client, bucketName string, likeService like.LikeContentService, saveService save.SaveContentService) *Handler {
	return &Handler{
		contentService: service,
		S3Client:       s3Client,
		BucketName:     bucketName,
		likeService:    likeService,
		saveService:    saveService,
	}
}

func (h *Handler) UploadContent(w http.ResponseWriter, r *http.Request) {

	const maxUploadSize = 15 << 20

	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)

	err := r.ParseMultipartForm(maxUploadSize)
	if err != nil {
		http.Error(w, "Maximum file size is 15MB", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")

	if err != nil {
		fmt.Fprintf(w, "Error retrieving the file: %v", err)
		return
	}

	uploaderIdStr := r.FormValue("uploaderId")

	if uploaderIdStr == "" {
		http.Error(w, "uploaderId must be filled", http.StatusBadRequest)
		return
	}

	uploaderId, err := strconv.ParseUint(uploaderIdStr, 10, 64)

	if err != nil {
		http.Error(w, "Invalid number parameter for 'uploaderId'", http.StatusBadRequest)
		return
	}

	caption := r.FormValue("caption")

	if caption == "" {
		http.Error(w, "caption must be filled", http.StatusBadRequest)
		return
	}

	fileType := handler.Header.Get("Content-Type")

	var contentType entity.ContentType

	if strings.HasPrefix(fileType, "image/") {
		contentType = "image"
	} else if strings.HasPrefix(fileType, "video/") {
		contentType = "video"
	} else {
		http.Error(w, "Unsupported file type", http.StatusBadRequest)
		return
	}

	newFileName := utils.GenerateFileName(handler)

	defer file.Close()

	fileUrl, err := utils.UploadFileToS3(h.S3Client, file, newFileName, h.BucketName, "Content")

	if err != nil {
		fmt.Fprintf(w, "Unable to upload file to S3: %v", err)
		return
	}

	content, err := h.contentService.UploadContent(uploaderId, caption, fileUrl, contentType)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(w, content)
}

func (h *Handler) UpdateContent(w http.ResponseWriter, r *http.Request) {
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

	content, err := h.contentService.UpdateContent(contentId, input.Caption)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(w, content)
}

func (h *Handler) GetAllContent(w http.ResponseWriter, r *http.Request) {
	var err error

	userId := utils.GetPathParam(r, "userId", "number", &err).(uint64)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	limit, err := strconv.Atoi(utils.GetQueryParam(r, "limit", "string", &err).(string))

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	cursor := utils.GetQueryParam(r, "cursor", "number", &err)

	var cursorPtr *uint64

	if cursor != nil {
		cursorValue := cursor.(uint64) 
		cursorPtr = &cursorValue    
	}

	contents, nextCursor, err := h.contentService.GetAllContents(limit, cursorPtr)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	var wg sync.WaitGroup
	contentList := make([]ContentResponse, len(*contents))

	likeChan := make(chan struct {
		index  int
		liked  bool
		likeId uint64
	}, len(*contents))

	saveChan := make(chan struct {
		index  int
		saved  bool
		saveId uint64
	}, len(*contents))

	for i, content := range *contents {
		wg.Add(2)

		go func(index int, contentId uint64) {
			defer wg.Done()

			like, _ := h.likeService.GetOneLike(userId, contentId)

			var likeId uint64
			if like != nil {
				likeId = like.Id
			} else {
				likeId = 0
			}

			likeChan <- struct {
				index  int
				liked  bool
				likeId uint64
			}{index: index, liked: like != nil, likeId: likeId}
		}(i, content.Id)

		go func(index int, contentId uint64) {
			defer wg.Done()

			save, _ := h.saveService.GetOneSave(userId, contentId)

			var saveId uint64
			if save != nil {
				saveId = save.Id
			} else {
				saveId = 0
			}

			saveChan <- struct {
				index  int
				saved  bool
				saveId uint64
			}{index: index, saved: save != nil, saveId: saveId}
		}(i, content.Id)
	}

	go func() {
		wg.Wait()
		close(likeChan)
		close(saveChan)
	}()

	for i := 0; i < len(*contents); i++ {
		likeResult := <-likeChan
		contentList[likeResult.index].Content = (*contents)[likeResult.index]
		contentList[likeResult.index].Like.IsLiked = likeResult.liked
		contentList[likeResult.index].Like.LikeId = likeResult.likeId
	}

	for i := 0; i < len(*contents); i++ {
		saveResult := <-saveChan
		contentList[saveResult.index].Content = (*contents)[saveResult.index]
		contentList[saveResult.index].Save.IsSaved = saveResult.saved
		contentList[saveResult.index].Save.SaveId = saveResult.saveId
	}

	response := ContentListResponse{
		Contents:   contentList,
		NextCursor: nextCursor,
	}

	utils.SuccessResponse(w, response)
}

func (h *Handler) GetOneContent(w http.ResponseWriter, r *http.Request) {
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

	content, err := h.contentService.GetOneContent(contentId)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	var response ContentResponse
	response.Content = *content

	var wg sync.WaitGroup

	likeChan := make(chan Like, 1)
	saveChan := make(chan Save, 1)

	wg.Add(1)
	go func() {
		defer wg.Done()
		like, _ := h.likeService.GetOneLike(userId, contentId)

		var likeId uint64
		if like != nil {
			likeId = like.Id
		} else {
			likeId = 0
		}

		likeChan <- Like{
			IsLiked: like != nil,
			LikeId:  likeId,
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		save, _ := h.saveService.GetOneSave(userId, contentId)

		var saveId uint64
		if save != nil {
			saveId = save.Id
		} else {
			saveId = 0
		}

		saveChan <- Save{
			IsSaved: save != nil,
			SaveId:  saveId,
		}
	}()

	go func() {
		wg.Wait()
		close(likeChan)
		close(saveChan)
	}()

	response.Like = <-likeChan
	response.Save = <-saveChan

	utils.SuccessResponse(w, response)
}

func (h *Handler) GetAllContentByFollowing(w http.ResponseWriter, r *http.Request) {
	var err error

	userId := utils.GetPathParam(r, "userId", "number", &err).(uint64)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	contents, err := h.contentService.GetAllContentsByFollowing(userId)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	var wg sync.WaitGroup
	response := make([]ContentResponse, len(*contents))

	likeChan := make(chan struct {
		index  int
		liked  bool
		likeId uint64
	}, len(*contents))

	saveChan := make(chan struct {
		index  int
		saved  bool
		saveId uint64
	}, len(*contents))

	for i, content := range *contents {
		wg.Add(2)

		go func(index int, contentId uint64) {
			defer wg.Done()

			like, _ := h.likeService.GetOneLike(userId, contentId)

			var likeId uint64
			if like != nil {
				likeId = like.Id
			} else {
				likeId = 0
			}

			likeChan <- struct {
				index  int
				liked  bool
				likeId uint64
			}{index: index, liked: like != nil, likeId: likeId}
		}(i, content.Id)

		go func(index int, contentId uint64) {
			defer wg.Done()

			save, _ := h.saveService.GetOneSave(userId, contentId)

			var saveId uint64
			if save != nil {
				saveId = save.Id
			} else {
				saveId = 0
			}

			saveChan <- struct {
				index  int
				saved  bool
				saveId uint64
			}{index: index, saved: save != nil, saveId: saveId}
		}(i, content.Id)
	}

	go func() {
		wg.Wait()
		close(likeChan)
		close(saveChan)
	}()

	for i := 0; i < len(*contents); i++ {
		likeResult := <-likeChan
		response[likeResult.index].Content = (*contents)[likeResult.index]
		response[likeResult.index].Like.IsLiked = likeResult.liked
		response[likeResult.index].Like.LikeId = likeResult.likeId
	}

	for i := 0; i < len(*contents); i++ {
		saveResult := <-saveChan
		response[saveResult.index].Content = (*contents)[saveResult.index]
		response[saveResult.index].Save.IsSaved = saveResult.saved
		response[saveResult.index].Save.SaveId = saveResult.saveId
	}

	utils.SuccessResponse(w, response)
}

func (h *Handler) DeleteContent(w http.ResponseWriter, r *http.Request) {
	var err error

	contentId := utils.GetPathParam(r, "contentId", "number", &err).(uint64)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	err = h.contentService.DeleteContent(contentId)

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
