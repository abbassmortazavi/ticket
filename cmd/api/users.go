package api

import (
	"net/http"
	"ticket/internal/store"
	"ticket/internal/utils"

	log2 "github.com/rs/zerolog/log"
)

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Mobile   string `json:"mobile"`
}

func (app *Application) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	err := utils.ReadJson(w, r, &req)
	if err != nil {
		log2.Err(err).Msg("error decoding body")
		utils.BadRequest(w, "error decoding body", err)
		return
	}

	ctx := r.Context()
	user := store.User{
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
