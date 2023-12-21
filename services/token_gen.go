package services

import (
	"time"

	db "study_marketplace/internal/database/queries"
	"study_marketplace/models"

	"github.com/golang-jwt/jwt"
)

var SecretKey = []byte("your-secret-key")

func generateToken(user db.User) (string, error) {
	claims := &models.Claims{
		UserID:   user.ID,
		Username: user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 10).Unix(), // Set token expiration time
			IssuedAt:  time.Now().Unix(),                          // Set token issued at time

		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
