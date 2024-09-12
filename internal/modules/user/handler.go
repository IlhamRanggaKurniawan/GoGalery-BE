package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/utils"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Handler struct {
	userService UserService
	S3Client    *s3.Client
	BucketName  string
}

type input struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	Bio             string `json:"bio"`
	ProfileUrl      string `json:"profileUrl"`
}

type authenticationRes struct {
	User        entity.User
	AccessToken string
}

func NewHandler(userService UserService, s3Client *s3.Client, bucketName string) Handler {
	return Handler{
		userService: userService,
		S3Client:    s3Client,
		BucketName:  bucketName,
	}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {

	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	if input.Password != input.ConfirmPassword {
		utils.ErrorResponse(w, fmt.Errorf("password doesn't match"), http.StatusBadRequest)
		return
	}

	user, err := h.userService.Register(input.Username, input.Email, input.Password)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	accessToken, err := utils.GenerateAccessToken(user.Username, user.Email, user.ID, user.Role, user.ProfileUrl, user.Bio)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	refreshToken, err := utils.GenerateRefreshToken(user.Username, user.Email, user.ID, user.Role, user.ProfileUrl, user.Bio)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "RefreshToken",
		Value:    refreshToken,
		Expires:  time.Now().Add(24 * time.Hour * 7),
		HttpOnly: true,
		Path:     "/",
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "AccessToken",
		Value:    accessToken,
		Expires:  time.Now().Add(5 * time.Minute),
		HttpOnly: true,
		Path:     "/",
	})

	user, err = h.userService.UpdateUser(user.ID, nil, nil, nil, &refreshToken)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	res := authenticationRes{
		User:        *user,
		AccessToken: accessToken,
	}

	utils.SuccessResponse(w, res)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {

	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	user, err := h.userService.Login(input.Username, input.Password)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	accessToken, err := utils.GenerateAccessToken(user.Username, user.Email, user.ID, user.Role, user.ProfileUrl, user.Bio)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	refreshToken, err := utils.GenerateRefreshToken(user.Username, user.Email, user.ID, user.Role, user.ProfileUrl, user.Bio)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "RefreshToken",
		Value:    refreshToken,
		Expires:  time.Now().Add(24 * time.Hour * 7),
		HttpOnly: true,
		Path:     "/",
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "AccessToken",
		Value:    accessToken,
		Expires:  time.Now().Add(5 * time.Minute),
		HttpOnly: true,
		Path:     "/",
	})

	user, err = h.userService.UpdateUser(user.ID, nil, nil, nil, &refreshToken)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	res := authenticationRes{
		User:        *user,
		AccessToken: accessToken,
	}

	utils.SuccessResponse(w, res)
}
func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	var err error

	id := utils.GetPathParam(r, "id", "number", &err).(uint64)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "RefreshToken",
		Value:    "",
		Expires:  time.Now().Add(-1),
		HttpOnly: true,
		Path:     "/",
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "AccessToken",
		Value:    "",
		Expires:  time.Now().Add(-1),
		HttpOnly: true,
		Path:     "/",
	})

	token := ""

	_, err = h.userService.UpdateUser(id, nil, nil, nil, &token)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	resp := struct {
		Message string `json:"message"`
	}{
		Message: "request success",
	}

	utils.SuccessResponse(w, resp)
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var err error

	userId := utils.GetPathParam(r, "id", "number", &err).(uint64)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	const maxUploadSize = 10 << 20

	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)

	err = r.ParseMultipartForm(maxUploadSize)

	if err != nil {
		utils.ErrorResponse(w, fmt.Errorf("maximum file size is 15MB"), http.StatusBadRequest)
		return
	}

	file, handler, _ := r.FormFile("file")

	bio := r.FormValue("bio")

	password := r.FormValue("password")

	profileUrl := r.FormValue("profileUrl")

	var url string

	if profileUrl == "" && file != nil {
		newFileName := utils.GenerateFileName(handler)

		url, err = utils.UploadFileToS3(h.S3Client, file, newFileName, h.BucketName, "Profile")
	} else if profileUrl != "" && file != nil {

		utils.UpdateFileInS3(h.S3Client, file, profileUrl, h.BucketName)
	}

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	user, err := h.userService.UpdateUser(userId, &bio, &url, &password, nil)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return

	}

	utils.SuccessResponse(w, user)
}

func (h *Handler) FindAllUsers(w http.ResponseWriter, r *http.Request) {
	var err error

	username := utils.GetPathParam(r, "username", "string", &err).(string)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	users, err := h.userService.FindAllUsersByUsername(username)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(w, users)
}

func (h *Handler) FindAllMutualUsers(w http.ResponseWriter, r *http.Request) {
	var err error

	userId := utils.GetPathParam(r, "userId", "number", &err).(uint64)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	users, err := h.userService.FindAllMutualUsers(userId)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(w, users)
}

func (h *Handler) FindUser(w http.ResponseWriter, r *http.Request) {
	var err error

	username := utils.GetPathParam(r, "username", "string", &err).(string)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	user, err := h.userService.FindOneUserByUsername(username)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(w, user)
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var err error

	userId := utils.GetPathParam(r, "userId", "number", &err).(uint64)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	err = h.userService.DeleteUser(userId)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "RefreshToken",
		Value:    "",
		Expires:  time.Now().Add(-1),
		HttpOnly: true,
		Path:     "/",
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "AccessToken",
		Value:    "",
		Expires:  time.Now().Add(-1),
		HttpOnly: true,
		Path:     "/",
	})

	resp := struct {
		Message string `json:"message"`
	}{
		Message: "request success",
	}

	utils.SuccessResponse(w, resp)
}

func (h *Handler) GetToken(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("RefreshToken")

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnauthorized)
		return
	}

	token, err := utils.ValidateToken(cookie.Value, "Refresh Token")

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnauthorized)
		return
	}

	accessToken, err := utils.GenerateAccessToken(token.Username, token.Email, token.ID, token.Role, nil, nil)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "AccessToken",
		Value:    accessToken,
		Expires:  time.Now().Add(5 * time.Minute),
		HttpOnly: true,
		Path:     "/",
	})

	resp := struct {
		AccessToken string `json:"accessToken"`
	}{
		AccessToken: accessToken,
	}

	utils.SuccessResponse(w, resp)
}
