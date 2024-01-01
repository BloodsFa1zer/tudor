package registry

import (
	"study_marketplace/pkg/domen/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func genTokFunc(secret string) func(userid int64, email string) (string, error) {
	return func(userid int64, userName string) (string, error) {
		claims := &models.AuthClaims{
			UserID: userid,
			Email:  userName,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 10)), // Set token expiration time
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte(secret))
		if err != nil {
			return "", err
		}
		return tokenString, nil
	}
}
