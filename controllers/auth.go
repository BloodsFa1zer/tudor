package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"study_marketplace/services"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

type AuthController interface {
	AuthWithProvider(ctx *gin.Context)
	AuthWithProviderCallback(ctx *gin.Context)
}

type authController struct {
	redirectPage string
	userService  services.UserService
}

func NewAuthController(redirectPage string, us services.UserService) AuthController {
	return &authController{redirectPage: redirectPage, userService: us}
}

func (t *authController) AuthWithProviderCallback(ctx *gin.Context) {
	provider := ctx.Param("provider")
	ctx.Set("provider", provider)
	fmt.Println("provider", provider)
	user, err := gothic.CompleteUserAuth(ctx.Writer, ctx.Request)
	if err != nil {
		ctx.AbortWithError(http.StatusForbidden, gin.Error{Err: errors.New("something went wrong")})
		return
	}
	fmt.Printf("%#v", user)
	ctx.Redirect(http.StatusFound, t.redirectPage)
}

func (t *authController) AuthWithProvider(ctx *gin.Context) {
	provider := ctx.Param("provider")
	ctx.Set("provider", provider)
	fmt.Println("provider", provider)
	gothic.BeginAuthHandler(ctx.Writer, ctx.Request)
}
