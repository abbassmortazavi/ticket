package models

import (
	"ticket/internal/modules/user/models"
	"time"
)

type Message struct {
	ID          int         `json:"id,omitempty"`
	RoomID      int         `json:"room_id,omitempty"`
	UserID      int         `json:"user_id,omitempty"`
	Message     string      `json:"message,omitempty"`
	MessageType string      `json:"message_type,omitempty"`
	CreatedAt   time.Time   `json:"created_at,omitempty"`
	UpdatedAt   time.Time   `json:"updated_at,omitempty"`
	User        models.User `json:"user,omitempty"`
	Room        Room        `json:"room,omitempty"`
}
