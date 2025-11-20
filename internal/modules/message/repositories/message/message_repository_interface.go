package message

import (
	"context"
	"ticket/internal/modules/message/models"
)

type MessageRepositoryInterface interface {
	SaveMessage(ctx context.Context, message *models.Message) (int, error)
	GetMessagesByRoomId(ctx context.Context, roomId int) ([]*models.Message, error)
}
