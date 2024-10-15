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

	user, err := utils.DecodeAccessToken(r)

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

	follow, err := h.followService.followUser(user.Id, input.FollowingId)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(w, follow)
}

func (h *Handler) GetAllFollows(w http.ResponseWriter, r *http.Request) {
	var err error

	user, err := utils.DecodeAccessToken(r)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	followerChan := make(chan *[]entity.Follow)
	followingChan := make(chan *[]entity.Follow)

	go func() {
		follower, _ := h.followService.GetAllFollower(user.Id)
		followerChan <- follower
	}()

	go func() {
		following, _ := h.followService.GetAllFollowing(user.Id)
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
	user, err := utils.DecodeAccessToken(r)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	followerChan := make(chan int64)
	followingChan := make(chan int64)

	go func() {
		follower, _ := h.followService.CountFollower(user.Id)
		followerChan <- follower
	}()

	go func() {
		following, _ := h.followService.CountFollowing(user.Id)
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
	user, err := utils.DecodeAccessToken(r)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	followingId := utils.GetQueryParam(r, "followingId", "number", &err).(uint64)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	follow, err := h.followService.CheckFollowing(user.Id, followingId)

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
