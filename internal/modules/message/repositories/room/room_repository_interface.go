package room

import (
	"context"
	"ticket/internal/modules/message/models"
)

type RoomRepositoryInterface interface {
	FindRoom(ctx context.Context, name string) (*models.Room, error)
}
