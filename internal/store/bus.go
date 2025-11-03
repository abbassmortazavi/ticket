package store

import (
	"context"
	"database/sql"
	"ticket/internal/models"

	"github.com/lib/pq"
)

type BusStore struct {
	db *sql.DB
}

func (s *BusStore) Create(ctx context.Context, bus models.Bus) (models.Bus, error) {
	q := `insert into buses (bus_number, operator_name,total_seats, bus_type,amenities)
		values ($1,$2,$3,$4,$5) returning id, bus_number, operator_name, total_seats, bus_type, amenities`
	err := s.db.QueryRowContext(ctx, q, bus.BusNumber, bus.OperatorName, bus.TotalSeats, bus.BusType, pq.Array(bus.Amenities)).Scan(
		&bus.ID,
		&bus.BusNumber,
		&bus.OperatorName,
		&bus.TotalSeats,
		&bus.BusType,
		pq.Array(&bus.Amenities),
	)
	if err != nil {
		return bus, err
	}
	return bus, nil
}
