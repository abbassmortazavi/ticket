package controllers

import (
	"net/http"
	"ticket/internal/modules/user/models"
	userService "ticket/internal/modules/user/services"
	"ticket/internal/utils"

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
