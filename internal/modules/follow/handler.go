package follow

import (
	"encoding/json"
	"net/http"

	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"
)

type Handler struct {
	followService FollowService
}

type input struct {
	ID              uint   `json:"id"`
	UserID          uint   `json:"userId"`
	FollowerID      uint   `json:"followerId"`
	FollowingID     uint   `json:"followingId"`
}

func NewHandler(followService FollowService) Handler {
	return Handler{followService}
}

func (h *Handler) FollowUser(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	follow, _ := h.followService.followUser(input.FollowerID, input.FollowingID)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(follow); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetAllFollows(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	follower, following, _ := h.followService.GetAllFollows(input.UserID)

	follow := struct{
		Follower *[]entity.Follow `json:"follower"`
		Following *[]entity.Follow `json:"following"`
	}{
		Follower: follower,
		Following: following,
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(follow); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) CheckFollowing(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	follow, _ := h.followService.CheckFollowing(input.FollowerID, input.FollowingID)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(follow); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) UnfollowUser(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.followService.UnfollowUser(input.ID)

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
