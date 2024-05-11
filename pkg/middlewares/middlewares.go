package middleware

import (
	"fmt"
	"net/http"
	"strings"
	reqm "study_marketplace/pkg/domain/mappers/reqresp_mappers"
	"study_marketplace/pkg/domain/models"
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

// CORS Allows server to use only secure(same-origin) resources to take data from
func (m *middleware) CORS() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cors := cors.New(
			cors.Config{
				// origins to take data from
				AllowOrigins: m.conf.AllowedOrigins,
				// methods that are secure to use
				AllowMethods: []string{"GET", "POST", "PATCH", "DELETE", "HEAD", "OPTIONS"},
				AllowHeaders: []string{"*"},
			})
		cors(ctx)
		// if the method is "OPTIONS" browser probably wants to check what is allowed by the developer
		if ctx.Request.Method == "OPTIONS" {
			ctx.Status(http.StatusOK)
			return
		}
		// if the method is not "OPTIONS" app just acts as if there is no any CORS-related checks
		ctx.Next()
	}

}

func (m *middleware) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authString := c.GetHeader("Authorization")
		if authString == "" {
			c.JSON(http.StatusUnauthorized, reqm.FailedResponse("Unauthorized"))
			c.Abort()
			return
		}
		authArray := strings.Split(authString, ":")
		bearerAndJwt := strings.Split(authArray[0], " ")
		if len(bearerAndJwt) != 2 && bearerAndJwt[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, reqm.FailedResponse("Unauthorized"))
			c.Abort()
			return
		}
		authJWT := bearerAndJwt[1]
		token, err := jwtValidate(authJWT, m.conf.JWTSecret)

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, reqm.FailedResponse("Unauthorized"))
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*models.AuthClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, reqm.FailedResponse("Unauthorized"))
			c.Abort()
			return
		}

		// You can access claims data here
		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)
		c.Next()
	}
}

func (m *middleware) PasswordMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authString := c.GetHeader("Authorization")
		if authString == "" {
			c.JSON(http.StatusUnauthorized, reqm.FailedResponse("Unauthorized"))
			c.Abort()
			return
		}
		// exppected token:pswd
		// pswdString := strings.Split(authString, ":")

		// if len(pswdString) != 2 {
		// 	// c.JSON(http.StatusUnauthorized, respmodels.FailedResponse{Data:"Not all info provided for change."))
		// 	// c.Abort()
		// 	// return
		// }

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
