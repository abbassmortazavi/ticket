package auth

import "log"

func GetJwtAuthenticator() *JWT {
	log.Println("Get JWT Authenticator", JwtAuthenticator)
	return JwtAuthenticator
}
