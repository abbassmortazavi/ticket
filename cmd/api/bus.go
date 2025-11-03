package api

import (
	"context"
	"net/http"
	"ticket/internal/models"
	"ticket/internal/utils"

	log2 "github.com/rs/zerolog/log"
)

type BusStoreRequest struct {
	BusNumber    string   `json:"bus_number"`
	OperatorName string   `json:"operator_name"`
	TotalSeats   int      `json:"total_seats"`
	BusType      string   `json:"bus_type"`
	Amenities    []string `json:"amenities"`
}

func (app *Application) CreateBus(w http.ResponseWriter, r *http.Request) {
	var req BusStoreRequest
	err := utils.ReadJson(w, r, &req)
	if err != nil {
		log2.Err(err).Msg("error decoding body")
		utils.BadRequest(w, "Validation Error!!", err)
		return
	}

	log2.Info().Interface("req => ", req).Msg("req")

	if !utils.ValidateStruct(w, &req) {
		return
	}

	ctx := context.Background()
	bus := models.Bus{
		BusNumber:    req.BusNumber,
		OperatorName: req.OperatorName,
		TotalSeats:   req.TotalSeats,
		BusType:      req.BusType,
		Amenities:    req.Amenities,
	}

	res, err := app.Store.Bus.Create(ctx, bus)
	if err != nil {
		log2.Err(err).Msg("error creating bus")
		utils.BadRequest(w, "Validation Error!!", err)
		return
	}

	utils.Created(w, res)

}
