package controllers

import (
	"log"
	"net/http"
	"ticket/internal/modules/auth/services"
	"ticket/internal/modules/user/models"
	userService "ticket/internal/modules/user/services"
	"ticket/internal/utils"
	"ticket/pkg/auth"

	log2 "github.com/rs/zerolog/log"
)

type Controller struct {
	authService services.AuthServiceInterface
	userService userService.UserServiceInterface
}

func New() *Controller {
	return &Controller{
		authService: services.New(auth.GetJwtAuthenticator()),
		userService: userService.New(),
	}
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type AuthResponse struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}

func (controller *Controller) Login(w http.ResponseWriter, r *http.Request) {
	log.Println("here", auth.GetJwtAuthenticator())
	var req LoginRequest
	err := utils.ReadJson(w, r, &req)
	if err != nil {
		log2.Err(err).Msg("error decoding body")
		utils.BadRequest(w, "error decoding body", err)
		return
	}
	if !utils.ValidateStruct(w, &req) {
		return
	}
	ctx := r.Context()
	user, err := controller.userService.GetByUsername(ctx, req.Username)
	if err != nil {
		utils.InternalError(w, err)
		return
	}
	// Check password
	if err := user.CheckPassword(req.Password); err == false {
		utils.InternalError(w, nil, "password error")
		return
	}
	token, err := controller.authService.GenerateToken(user.ID, user.Username)
	if err != nil {
		utils.InternalError(w, err)
		return
	}
	res := AuthResponse{
		Token: token,
		User:  user,
	}
	utils.Success(w, http.StatusOK, res, "User Login")
}

func (controller *Controller) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := utils.ReadJson(w, r, &user)
	if err != nil {
		log2.Err(err).Msg("error decoding body")
		utils.BadRequest(w, "error decoding body", err)
		return
	}
	if !utils.ValidateStruct(w, &user) {
		return
	}
	ctx := r.Context()
	if err := user.HashPassword(user.Password); err != nil {
		utils.InternalError(w, err)
		return
	}

	res, err := controller.userService.CreateUser(ctx, user)
	if err != nil {
		utils.InternalError(w, err)
		return
	}
	utils.Created(w, res)
}
