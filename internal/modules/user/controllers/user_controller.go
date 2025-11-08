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
