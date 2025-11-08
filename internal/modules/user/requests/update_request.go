package requests

type UpdateRequest struct {
	ID       int    `json:"_"`
	Username string `json:"username" validate:"omitempty,min=3,max=32,alphanum"`
	Password string `json:"password" validate:"omitempty,min=6"`
	Email    string `json:"email" validate:"omitempty,email"`
	FullName string `json:"full_name" validate:"omitempty,max=255"`
	Mobile   string `json:"mobile" validate:"omitempty"`
}
