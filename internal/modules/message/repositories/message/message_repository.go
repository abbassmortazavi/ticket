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
func (m *MessageRepository) GetMessagesByRoomId(ctx context.Context, roomId int) ([]*models.Message, error) {
	query := `select id, room_id, user_id, message, message_type, created_at from messages where room_id = $1 order by id desc`
	rows, err := m.DB.QueryContext(ctx, query, roomId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var messages []*models.Message
	for rows.Next() {
		var msg models.Message
		err := rows.Scan(&msg.ID, &msg.RoomID, &msg.UserID, &msg.Message, &msg.MessageType, &msg.CreatedAt)
		if err != nil {
			return nil, err
		}
		messages = append(messages, &msg)
	}
	return messages, nil
}
