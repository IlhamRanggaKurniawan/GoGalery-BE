package content

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/utils"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Handler struct {
	contentService ContentService
	S3Client       *s3.Client
	BucketName     string
}

type input struct {
	ID         uint64 `json:"id"`
	UploaderID uint64 `json:"uploaderId"`
	UserID     uint64 `json:"userId"`
	Caption    string `json:"caption"`
	Path       string `json:"path"`
}

func NewHandler(service ContentService, s3Client *s3.Client, bucketName string) *Handler {
	return &Handler{
		contentService: service,
		S3Client:       s3Client,
		BucketName:     bucketName,
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
	
	fmt.Println(file)
	if err != nil {
		fmt.Fprintf(w, "Error retrieving the file: %v", err)
		return
	}

	uploaderIdStr := r.FormValue("uploaderId")
	fmt.Println(uploaderIdStr)


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

	fileUrl, err := utils.UploadFileToS3(h.S3Client, file, newFileName, h.BucketName)

	if err != nil {
		fmt.Fprintf(w, "Unable to upload file to S3: %v", err)
		return
	}

	content, _ := h.contentService.UploadContent(uploaderId, caption, fileUrl, contentType)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(content); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// func (h *Handler) UploadContent(w http.ResponseWriter, r *http.Request) {
// 	var input input

// 	err := json.NewDecoder(r.Body).Decode(&input)

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	content, _ := h.contentService.UploadContent(input.UploaderID, input.Caption, input.Path)

// 	w.Header().Set("Content-Type", "application/json")

// 	if err := json.NewEncoder(w).Encode(content); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }

func (h *Handler) UpdateContent(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	content, _ := h.contentService.UpdateContent(input.ID, input.Caption)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(content); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetAllContent(w http.ResponseWriter, r *http.Request) {

	contents, _ := h.contentService.GetAllContents()

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(contents); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetOneContent(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	content, _ := h.contentService.GetOneContent(input.ID)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(content); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetAllContentByFollowing(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	content, _ := h.contentService.GetAllContentsByFollowing(input.UserID)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(content); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DeleteContent(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.contentService.DeleteContent(input.ID)

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
