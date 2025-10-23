package api

import (
	"net/http"
	"strconv"
	"ticket/internal/models"
	"ticket/internal/utils"

	"github.com/go-chi/chi/v5"
	log2 "github.com/rs/zerolog/log"
)

type CreateUserRequest struct {
	Username string `json:"username" validate:"required,min=3,max=32,alphanum"`
	Password string `json:"password" validate:"required,min=6"`
	Email    string `json:"email" validate:"required,email"`
	FullName string `json:"full_name" validate:"max=255"`
	Mobile   string `json:"mobile" validate:"required"`
}

func (app *Application) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
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
	user := models.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		FullName: req.FullName,
		Mobile:   req.Mobile,
	}
	res, err := app.Store.User.Create(ctx, user)
	if err != nil {
		utils.InternalError(w, err)
		return
	}
	utils.Created(w, res)
}
func (app *Application) GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	ctx := r.Context()
	userID, err := strconv.Atoi(id)
	if err != nil {
		log2.Err(err).Msg("error decoding id")
		utils.BadRequest(w, "error decoding id", err)
	}
	user, err := app.Store.User.GetUser(ctx, userID)
	if err != nil {
		utils.InternalError(w, err)
		return
	}

	res := models.User{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		FullName: user.FullName,
		Mobile:   user.Mobile,
	}
	utils.Success(w, http.StatusOK, res, "Fetch User Successfully")

}
func (app *Application) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	ctx := r.Context()
	userID, err := strconv.Atoi(id)
	if err != nil {
		log2.Err(err).Msg("error decoding id")
		utils.BadRequest(w, "error decoding id", err)
		return
	}
	err = app.Store.User.Delete(ctx, userID)
	if err != nil {
		utils.InternalError(w, err)
		return
	}

	utils.Success(w, http.StatusOK, nil, "User Delete Successfully!!")
}

type UpdateRequest struct {
	ID       int    `json:"_"`
	Username string `json:"username" validate:"omitempty,min=3,max=32,alphanum"`
	Password string `json:"password" validate:"omitempty,min=6"`
	Email    string `json:"email" validate:"omitempty,email"`
	FullName string `json:"full_name" validate:"omitempty,max=255"`
	Mobile   string `json:"mobile" validate:"omitempty"`
}

func (app *Application) Update(w http.ResponseWriter, r *http.Request) {
	var req UpdateRequest
	err := utils.ReadJson(w, r, &req)
	if err != nil {
		log2.Err(err).Msg("error decoding body")
		utils.BadRequest(w, "error decoding body", err)
		return
	}

	if !utils.ValidateStruct(w, &req) {
		return
	}

	id := chi.URLParam(r, "id")
	ctx := r.Context()
	userID, err := strconv.Atoi(id)
	if err != nil {
		log2.Err(err).Msg("error decoding id")
		utils.BadRequest(w, "error decoding id", err)
		return
	}
	user, err := app.Store.User.GetUser(ctx, userID)
	if err != nil {
		utils.InternalError(w, err, "User Not Found!!")
		return
	}
	if req.Username != "" {
		user.Username = req.Username
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.FullName != "" {
		user.FullName = req.FullName
	}
	if req.Mobile != "" {
		user.Mobile = req.Mobile
	}
	res, err := app.Store.User.Update(ctx, user)
	log2.Err(err).Msg("error updating user")
	if err != nil {
		utils.InternalError(w, err)
		return
	}
	utils.Success(w, http.StatusOK, res, "User Update Successfully")
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type AuthResponse struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}

func (app *Application) Login(w http.ResponseWriter, r *http.Request) {
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
	user, err := app.Store.User.GetUserByUsername(ctx, req.Username)
	if err != nil {
		utils.InternalError(w, err)
		return
	}
	// Check password
	if err := user.CheckPassword(req.Password); err == false {
		utils.InternalError(w, nil, "password error")
		return
	}
	token, err := app.Authenticator.GenerateToken(user.ID, user.Username)
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
func (app *Application) Register(w http.ResponseWriter, r *http.Request) {
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
	res, err := app.Store.User.Create(ctx, user)
	if err != nil {
		utils.InternalError(w, err)
		return
	}
	utils.Created(w, res)
}
