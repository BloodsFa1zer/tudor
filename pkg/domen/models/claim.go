package models

import (
	"github.com/golang-jwt/jwt/v5"
)

type AuthClaims struct {
	UserID int64  `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}
