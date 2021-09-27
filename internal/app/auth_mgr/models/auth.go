package models

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)


// Claims Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string `json:"username"`
	UserID string `json:"user_id"`
	Role string `json:"role"`
	jwt.StandardClaims
}

type Token struct {
	Type string `json:"type"`
	Value string `json:"value"`
	ExpiresAt time.Time `json:"expires_at"`
}
