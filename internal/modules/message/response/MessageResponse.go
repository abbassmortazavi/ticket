package response

import "time"

type MessageResponse struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Username    string    `json:"username"`
	Message     string    `json:"message"`
	MessageType string    `json:"message_type"`
	IsOwn       bool      `json:"is_own"`
	CreatedAt   time.Time `json:"created_at"`
}
