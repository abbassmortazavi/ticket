package models

import (
	"ticket/internal/modules/user/models"
	"time"
)

type Room struct {
	ID          string      `json:"id,omitempty"`
	RoomName    string      `json:"room_name,omitempty"`
	CreatedBy   string      `json:"user_id,omitempty"`
	Description string      `json:"message,omitempty"`
	MessageType int         `json:"message_type,omitempty"`
	IsPublic    int         `json:"is_public,omitempty"`
	IsActive    int         `json:"is_active,omitempty"`
	MaxUsers    int         `json:"max_users,omitempty"`
	CreatedAt   time.Time   `json:"created_at,omitempty"`
	UpdatedAt   time.Time   `json:"updated_at,omitempty"`
	User        models.User `json:"user,omitempty"`
}
