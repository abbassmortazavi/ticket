package auth

type Authenticator interface {
	GenerateToken(userID int, username string) (string, error)
	ValidateToken(token string) (*Claims, error)
}
