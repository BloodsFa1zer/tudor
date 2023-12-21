package middleware

import (
	"net/http"
	"strings"
	config "study_marketplace/config"
	"study_marketplace/controllers"
	"study_marketplace/models"
	"study_marketplace/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Middleware interface {
	CORS() gin.HandlerFunc
	AuthMiddleware() gin.HandlerFunc
	PasswordMiddleware(controller controllers.UserController) gin.HandlerFunc
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
			AllowOrigins: m.conf.AllowedOrigins,
			AllowMethods: []string{"GET", "POST", "PATCH", "DELETE", "HEAD", "OPTIONS"},
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
			c.JSON(http.StatusUnauthorized, models.NewResponseFailed("Unauthorized"))
			c.Abort()
			return
		}
		authArray := strings.Split(authString, ":")
		authJWT := authArray[0]

		token, err := jwt.ParseWithClaims(authJWT, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return services.SecretKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, models.NewResponseFailed("Unauthorized"))
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*models.Claims)
		if !ok {
			c.JSON(http.StatusUnauthorized, models.NewResponseFailed("Unautorized"))
			c.Abort()
			return
		}

		// You can access claims data here
		c.Set("user_id", claims.UserID)
		c.Next()
	}
}

func (m *middleware) PasswordMiddleware(controller controllers.UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		authString := c.GetHeader("Authorization")
		if authString == "" {
			c.JSON(http.StatusUnauthorized, models.NewResponseFailed("Unauthorized"))
			c.Abort()
			return
		}
		// exppected token:pswd
		pswdString := strings.Split(authString, ":")

		if len(pswdString) != 2 {
			c.JSON(http.StatusUnauthorized, models.NewResponseFailed("Not all info provided for change."))
			c.Abort()
			return
		}

		pswd := pswdString[1]
		userPswd := controller.GetPassword(c)
		err := services.ComparePassword(userPswd, pswd)

		if err != nil {
			c.JSON(http.StatusUnauthorized, models.NewResponseFailed("Unauthorized"))
			c.Abort()
			return
		}

		c.Next()
	}
}
