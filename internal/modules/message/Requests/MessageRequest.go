package Requests

type MessageRequest struct {
	RoomID      int    `json:"room_id" validate:"required"`
	UserID      int    `json:"user_id,omitempty" validate:"required"`
	Message     string `json:"message,omitempty" validate:"required"`
	MessageType string `json:"message_type,omitempty" validate:"required"`
}
