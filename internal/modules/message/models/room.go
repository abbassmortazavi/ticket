package models

import (
	"ticket/internal/modules/user/models"
	"time"
)

type Room struct {
	ID          int         `json:"id,omitempty"`
	RoomName    string      `json:"room_name,omitempty"`
	Description string      `json:"message,omitempty"`
	CreatedBy   int         `json:"user_id,omitempty"`
	MessageType string      `json:"message_type,omitempty"`
	IsPublic    bool        `json:"is_public,omitempty"`
	IsActive    bool        `json:"is_active,omitempty"`
	MaxUsers    int         `json:"max_users,omitempty"`
	CreatedAt   time.Time   `json:"created_at,omitempty"`
	UpdatedAt   time.Time   `json:"updated_at,omitempty"`
	User        models.User `json:"user,omitempty"`
}
