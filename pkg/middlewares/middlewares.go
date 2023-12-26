package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"study_marketplace/pkg/domen/models"
	respmodels "study_marketplace/pkg/domen/models/response_models"
	config "study_marketplace/pkg/infrastructure/config"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Middleware interface {
	CORS() gin.HandlerFunc
	AuthMiddleware() gin.HandlerFunc
	PasswordMiddleware() gin.HandlerFunc
}

type middleware struct {
	conf *config.Config
}

func NewMiddleware(conf *config.Config) Middleware {
	return &middleware{conf}
}

func (m *middleware) CORS() gin.HandlerFunc {
	cors := cors.New(
		cors.Config{
			AllowWildcard: true,
			AllowOrigins:  m.conf.AllowedOrigins,
			AllowMethods:  []string{"GET", "POST", "PATCH", "DELETE", "HEAD", "OPTIONS"},
			AllowHeaders: []string{
				"Origin", "Content-Length", "Content-Type",
				"Access-Control-Allow-Headers", "Access-Control-Request-Method",
				"Access-Control-Request-Headers", "Access-Control-Allow-Origin",
				"X-Requested-With", "Accept", "Authorization"},
		})

	return cors
}

func (m *middleware) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authString := c.GetHeader("Authorization")
		if authString == "" {
			c.JSON(http.StatusUnauthorized, respmodels.NewResponseFailed("Unauthorized"))
			c.Abort()
			return
		}
		authArray := strings.Split(authString, ":")
		authJWT := authArray[0]

		token, err := jwtValidate(authJWT, m.conf.JWTSecret)

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, respmodels.NewResponseFailed("Unauthorized"))
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*models.AuthClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, respmodels.NewResponseFailed("Unautorized"))
			c.Abort()
			return
		}

		// You can access claims data here
		c.Set("user_id", claims.UserID)
		c.Next()
	}
}

func (m *middleware) PasswordMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authString := c.GetHeader("Authorization")
		if authString == "" {
			c.JSON(http.StatusUnauthorized, respmodels.NewResponseFailed("Unauthorized"))
			c.Abort()
			return
		}
		// exppected token:pswd
		pswdString := strings.Split(authString, ":")

		if len(pswdString) != 2 {
			c.JSON(http.StatusUnauthorized, respmodels.NewResponseFailed("Not all info provided for change."))
			c.Abort()
			return
		}

		c.Next()
	}
}

func jwtValidate(token, signedStr string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &models.AuthClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(signedStr), nil
	})
}
