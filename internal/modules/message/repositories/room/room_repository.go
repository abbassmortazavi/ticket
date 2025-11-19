package room

import (
	"context"
	"database/sql"
	"ticket/internal/modules/message/models"
	"ticket/pkg/database"
)

type RoomRepository struct {
	DB *sql.DB
}

func New() *RoomRepository {
	return &RoomRepository{
		DB: database.Connection(),
	}
}

func (r *RoomRepository) FindRoom(ctx context.Context, name string) (*models.Room, error) {
	var room models.Room
	query := `SELECT * FROM rooms WHERE room_name = $1;`
	row := r.DB.QueryRowContext(ctx, query, name)
	err := row.Scan(
		&room.ID,
		&room.RoomName,
		&room.Description,
		&room.CreatedBy,
		&room.IsPublic,
		&room.IsActive,
		&room.MaxUsers,
		&room.CreatedAt,
		&room.UpdatedAt)
	return &room, err
}
