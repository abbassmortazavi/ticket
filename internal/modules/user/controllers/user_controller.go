package controllers

import (
	"net/http"
	"strconv"
	"ticket/internal/modules/user/models"
	userService "ticket/internal/modules/user/services"
	"ticket/internal/utils"

	"github.com/go-chi/chi/v5"
	log2 "github.com/rs/zerolog/log"
)

type UserController struct {
	userService userService.UserServiceInterface
}

func New() *UserController {
	return &UserController{
		userService: userService.New(),
	}
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
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

	res, err := c.userService.CreateUser(ctx, user)
	if err != nil {
		utils.InternalError(w, err)
		return
	}
	utils.Created(w, res)
}
func (c *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		log2.Err(err).Msg("error decoding id")
		utils.BadRequest(w, "error decoding id", err)
		return
	}
	ctx := r.Context()
	user, err := c.userService.GetById(ctx, userID)
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
func (c *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		log2.Err(err).Msg("error decoding id")
		utils.BadRequest(w, "error decoding id", err)
		return
	}
	ctx := r.Context()
	user, err := c.userService.GetById(ctx, userID)
	if err != nil {
		utils.InternalError(w, err)
		return
	}
	err = c.userService.Delete(ctx, user.ID)
	if err != nil {
		utils.InternalError(w, err)
		return
	}
	utils.Success(w, http.StatusOK, nil, "User Deleted Successfully")
}

type UpdateRequest struct {
	ID       int    `json:"_"`
	Username string `json:"username" validate:"omitempty,min=3,max=32,alphanum"`
	Password string `json:"password" validate:"omitempty,min=6"`
	Email    string `json:"email" validate:"omitempty,email"`
	FullName string `json:"full_name" validate:"omitempty,max=255"`
	Mobile   string `json:"mobile" validate:"omitempty"`
}

func (c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		log2.Err(err).Msg("error decoding id")
		utils.BadRequest(w, "error decoding id", err)
		return
	}
	var req UpdateRequest
	err = utils.ReadJson(w, r, &req)
	if err != nil {
		log2.Err(err).Msg("error decoding body")
		utils.BadRequest(w, "error decoding body", err)
		return
	}
	if !utils.ValidateStruct(w, &req) {
		return
	}
	ctx := r.Context()
	user, err := c.userService.GetById(ctx, userID)
	if err != nil {
		utils.InternalError(w, err)
		return
	}
	if err := user.HashPassword(req.Password); err != nil {
		utils.InternalError(w, err)
		return
	}
	data := models.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		FullName: req.FullName,
		Mobile:   req.Mobile,
	}
	res := c.userService.Update(ctx, data)
	utils.Success(w, http.StatusOK, res, "Update User Successfully")
}
