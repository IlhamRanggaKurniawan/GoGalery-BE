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
	Id              uint64 `json:"id"`
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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if input.Password != input.ConfirmPassword {
		http.Error(w, "Password doen't match", http.StatusBadRequest)
		return
	}

	user, _ := h.userService.Register(input.Username, input.Email, input.Password)

	accessToken, err := utils.GenerateAccessToken(user.Username, user.Email, user.ID, user.Role, user.ProfileUrl, user.Bio)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	refreshToken, err := utils.GenerateRefreshToken(user.Username, user.Email, user.ID, user.Role, user.ProfileUrl, user.Bio)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	res := authenticationRes{
		User:        *user,
		AccessToken: accessToken,
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {

	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.userService.Login(input.Username, input.Password)

	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusBadRequest)
		return
	}

	accessToken, err := utils.GenerateAccessToken(user.Username, user.Email, user.ID, user.Role, user.ProfileUrl, user.Bio)

	if err != nil {
		http.Error(w, "Failed to generate access token", http.StatusInternalServerError)
		return
	}

	refreshToken, err := utils.GenerateRefreshToken(user.Username, user.Email, user.ID, user.Role, user.ProfileUrl, user.Bio)
	if err != nil {
		http.Error(w, "Failed to generate refresh token", http.StatusInternalServerError)
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
		http.Error(w, "Failed to update user with refresh token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	res := authenticationRes{
		User:        *user,
		AccessToken: accessToken,
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	id := utils.GetPathParam(w, r, "id", "number").(uint64)

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

	_, err := h.userService.UpdateUser(id, nil, nil, nil, &token)

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

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userId := utils.GetPathParam(w, r, "id", "number").(uint64)

	const maxUploadSize = 10 << 20

	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)

	err := r.ParseMultipartForm(maxUploadSize)

	if err != nil {
		http.Error(w, "Maximum file size is 15MB", http.StatusBadRequest)
		return
	}

	file, handler, _ := r.FormFile("file")

	bio := r.FormValue("bio")

	password := r.FormValue("password")

	fmt.Println(password)

	profileUrl := r.FormValue("profileUrl")

	fmt.Println(profileUrl)

	var url string

	if profileUrl == "" && file != nil {
		newFileName := utils.GenerateFileName(handler)

		url, err = utils.UploadFileToS3(h.S3Client, file, newFileName, h.BucketName, "Profile")
	} else if profileUrl != "" && file != nil {

		utils.UpdateFileInS3(h.S3Client, file, profileUrl, h.BucketName)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("Tes")

	user, err := h.userService.UpdateUser(userId, &bio, &url, &password, nil)

	fmt.Println("after")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(user)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) FindAllUsers(w http.ResponseWriter, r *http.Request) {
	username := utils.GetPathParam(w, r, "username", "string").(string)

	users, _ := h.userService.FindAllUsers(username)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) FindAllMutualUsers(w http.ResponseWriter, r *http.Request) {
	userId := utils.GetPathParam(w, r, "userId", "number").(uint64)

	users, _ := h.userService.FindAllMutualUsers(userId)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) FindUser(w http.ResponseWriter, r *http.Request) {

	username := utils.GetPathParam(w, r, "username", "string").(string)

	if username == "" {
		http.Error(w, "params is empty", http.StatusBadRequest)
		return
	}

	user, err := h.userService.FindOneUser(username)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {

	userId := utils.GetPathParam(w, r, "userId", "number").(uint64)

	err := h.userService.DeleteUser(userId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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

func (h *Handler) GetToken(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	cookie, err := r.Cookie("RefreshToken")

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Refresh token not found"})
		return
	}

	token, err := utils.ValidateToken(cookie.Value, "Refresh Token")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid refresh token"})
		return
	}

	accessToken, err := utils.GenerateAccessToken(token.Username, token.Email, token.ID, token.Role, nil, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to generate access token"})
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

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
}
