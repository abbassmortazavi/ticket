package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"ticket/internal/modules/auth/requests"
	"ticket/internal/modules/auth/services"
	"ticket/internal/modules/user/models"
	userService "ticket/internal/modules/user/services"
	"ticket/internal/utils"
	"ticket/pkg/auth"
	"ticket/pkg/html"
	"ticket/pkg/session"

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

type AuthResponse struct {
	Tokens *auth.TokenResponse `json:"tokens"`
	User   models.User         `json:"user"`
}

func (controller *Controller) Login(w http.ResponseWriter, r *http.Request) {
	var req requests.LoginRequest
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
	tokens := controller.authService.GenerateToken(user.ID, user.Username)

	res := AuthResponse{
		Tokens: tokens,
		User:   user,
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
func (controller *Controller) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var request struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Refresh access token
	tokens, err := controller.authService.RefreshToken(request.RefreshToken)
	if err != nil {
		http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
		return
	}

	res := AuthResponse{
		Tokens: tokens,
	}
	utils.Success(w, http.StatusOK, res, "Refresh Token Success")
}

// Web page
func (controller *Controller) LoginForm(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"title": "Home Page",
	}
	html.Render(w, "login", data)
}
func (controller *Controller) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req requests.LoginRequest
	err := utils.ReadJson(w, r, &req)
	if err != nil {
		log.Println(1)
		log2.Err(err).Msg("error decoding body")
		return
	}
	log.Println(2)
	if !utils.ValidateStruct(w, &req) {
		return
	}
	ctx := r.Context()
	user, err := controller.userService.GetByUsername(ctx, req.Username)
	if err != nil {
		utils.InternalError(w, err)
		return
	}
	if err := user.CheckPassword(req.Password); err == false {
		utils.InternalError(w, nil, "password error")
		return
	}

	err = session.Set(w, r, "user_id", user.ID)
	if err != nil {
		log2.Err(err).Msg("error setting session")
		return
	}

	err = session.Set(w, r, "name", user.FullName)
	if err != nil {
		log2.Err(err).Msg("error setting session")
		return
	}
	http.Redirect(w, r, "/join", http.StatusFound)
}

func (c *Controller) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	userID, err := session.Get(r, "user_id")
	if err != nil {
		err = utils.WriteJsonError(w, http.StatusUnauthorized, "Please log in")
		if err != nil {
			return
		}
		return
	}

	userName, err := session.Get(r, "name")
	if err != nil {
		err = utils.WriteJsonError(w, http.StatusInternalServerError, "Session error")
		if err != nil {
			return
		}
		return
	}

	err = utils.WriteJson(w, http.StatusOK, map[string]interface{}{
		"success":  true,
		"user_id":  userID,
		"username": userName,
	})
	if err != nil {
		return
	}
}
