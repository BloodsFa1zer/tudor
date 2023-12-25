package controllers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	reqresmappers "study_marketplace/pkg/domen/mappers/req_res_mappers"
	"study_marketplace/pkg/services"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

type AuthController interface {
	AuthWithProvider(ctx *gin.Context)
	AuthWithProviderCallback(ctx *gin.Context)
}

type authController struct {
	redirectPage string
	services.UserService
}

func NewAuthController(redirectPage string, us services.UserService) AuthController {
	return &authController{redirectPage: redirectPage, UserService: us}
}

func (c *authController) AuthWithProviderCallback(ctx *gin.Context) {
	provider := ctx.Param("provider")
	ctx.Request = ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), "provider", provider))
	fmt.Println("provider", provider)
	user, err := gothic.CompleteUserAuth(ctx.Writer, ctx.Request)
	if err != nil {
		ctx.AbortWithError(http.StatusForbidden, gin.Error{Err: errors.New("something went wrong")})
		return
	}
	token, err := c.ProviderAuth(ctx, reqresmappers.GothToUserToUser(user))
	if err != nil {
		ctx.AbortWithError(http.StatusForbidden, gin.Error{Err: errors.New("something went wrong")})
		return
	}
	ctx.Request.Header.Add("Authorization", token)
	ctx.AddParam("token", token) //TODO: get away token from params
	ctx.Redirect(http.StatusPermanentRedirect, c.redirectPage+"/:token")
}

// @Auth-with-provider			godoc
// @Summary						GET request for auth with provider
// @Description					requires param provider, for example google, facebook or apple  (at this moment apple not working) This request redirects to the provider's page for authorization, which in turn transmits a token in the parameters (token) and header (Authorization)
// @Tags						auth_with_provider get request for auth with provider
// @Accept						html
// @Produce						html
// @Param						provider	path		string		true	"provider for auth"
// @Success 					302
// @Router						/api/auth/{provider} [get]
func (c *authController) AuthWithProvider(ctx *gin.Context) {
	provider := ctx.Param("provider")
	ctx.Request = ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), "provider", provider))
	gothic.BeginAuthHandler(ctx.Writer, ctx.Request)
}
