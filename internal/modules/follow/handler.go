package follow

import (
	"encoding/json"
	"net/http"

	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/database/entity"
	"github.com/IlhamRanggaKurniawan/ConnectVerse-BE/internal/utils"
)

type Handler struct {
	followService FollowService
}

type input struct {
	UserId      uint64 `json:"userId"`
	FollowerId  uint64 `json:"followerId"`
	FollowingId uint64 `json:"followingId"`
}

func NewHandler(followService FollowService) Handler {
	return Handler{followService}
}

func (h *Handler) FollowUser(w http.ResponseWriter, r *http.Request) {
	var input input

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	follow, err := h.followService.followUser(input.FollowerId, input.FollowingId)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(w, follow)
}

func (h *Handler) GetAllFollows(w http.ResponseWriter, r *http.Request) {
	var err error

	userId := utils.GetPathParam(r, "userId", "number", &err).(uint64)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	followerChan := make(chan *[]entity.Follow)
	followingChan := make(chan *[]entity.Follow)

	go func() {
		follower, _ := h.followService.GetAllFollower(userId)
		followerChan <- follower
	}()

	go func() {
		following, _ := h.followService.GetAllFollowing(userId)
		followingChan <- following
	}()

	follow := struct {
		Follower  *[]entity.Follow `json:"follower"`
		Following *[]entity.Follow `json:"following"`
	}{
		Follower:  <-followerChan,
		Following: <-followingChan,
	}

	utils.SuccessResponse(w, follow)
}

func (h *Handler) CountFollow(w http.ResponseWriter, r *http.Request) {
	var err error

	userId := utils.GetPathParam(r, "userId", "number", &err).(uint64)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	followerChan := make(chan int64)
	followingChan := make(chan int64)

	go func() {
		follower, _ := h.followService.CountFollower(userId)
		followerChan <- follower
	}()

	go func() {
		following, _ := h.followService.CountFollowing(userId)
		followingChan <- following
	}()

	follow := struct {
		Follower  int64 `json:"follower"`
		Following int64 `json:"following"`
	}{
		Follower:  <-followerChan,
		Following: <-followingChan,
	}

	utils.SuccessResponse(w, follow)
}

func (h *Handler) CheckFollowing(w http.ResponseWriter, r *http.Request) {
	var err error

	followerId := utils.GetQueryParam(r, "followerId", "number", &err).(uint64)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	followingId := utils.GetQueryParam(r, "followingId", "number", &err).(uint64)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	follow, err := h.followService.CheckFollowing(followerId, followingId)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(w, follow)
}

func (h *Handler) UnfollowUser(w http.ResponseWriter, r *http.Request) {
	var err error

	followId := utils.GetPathParam(r, "followId", "number", &err).(uint64)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	err = h.followService.UnfollowUser(followId)

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
