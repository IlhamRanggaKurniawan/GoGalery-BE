package user

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/utils"
)

type Handler struct {
	userService UserService
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

func NewHandler(userService UserService) Handler {
	return Handler{userService}
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

	accessToken, err := utils.GenerateAccessToken(user.Username, user.Email, user.ID, user.Role)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	refreshToken, err := utils.GenerateRefreshToken(user.Username, user.Email, user.ID, user.Role)

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

	user, err = h.userService.UpdateUser(user.Username, nil, nil, nil, &refreshToken)

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

	accessToken, err := utils.GenerateAccessToken(user.Username, user.Email, user.ID, user.Role)
	if err != nil {
		http.Error(w, "Failed to generate access token", http.StatusInternalServerError)
		return
	}

	refreshToken, err := utils.GenerateRefreshToken(user.Username, user.Email, user.ID, user.Role)
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

	user, err = h.userService.UpdateUser(user.Username, nil, nil, nil, &refreshToken)
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
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

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

	token := ""

	_, err = h.userService.UpdateUser(input.Username, nil, nil, nil, &token)

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
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, _ := h.userService.UpdateUser(input.Username, &input.Bio, &input.ProfileUrl, &input.Password, nil)

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
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.userService.DeleteUser(input.Id)

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

	accessToken, err := utils.GenerateAccessToken(token.Username, token.Email, token.ID, token.Role)
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
