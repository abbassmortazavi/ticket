package message

import (
	"context"
	"database/sql"
	"ticket/internal/modules/message/models"
	"ticket/pkg/database"
	"time"
)

type MessageRepository struct {
	DB *sql.DB
}

func New() *MessageRepository {
	return &MessageRepository{
		DB: database.Connection(),
	}
}
func (m *MessageRepository) SaveMessage(ctx context.Context, message *models.Message) (int, error) {
	query := `INSERT INTO messages(room_id, user_id, message, message_type) VALUES ($1, $2, $3, $4) returning id`
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, query, message.RoomID, message.UserID, message.Message, message.MessageType).Scan(
		&message.ID)
	if err != nil {
		return 0, err
	}
	return message.ID, err
}
