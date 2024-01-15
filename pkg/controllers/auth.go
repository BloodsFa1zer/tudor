package controllers

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"study_marketplace/pkg/domain/models/entities"
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
	callbackfunc func(res http.ResponseWriter, req *http.Request) (*entities.User, error)
	services.AuthService
}

func NewAuthController(redirectPage string,
	callbackfunc func(res http.ResponseWriter, req *http.Request) (*entities.User, error),
	us services.AuthService) AuthController {
	return &authController{redirectPage: redirectPage, callbackfunc: callbackfunc, AuthService: us}
}

func (c *authController) AuthWithProviderCallback(ctx *gin.Context) {
	provider := ctx.Param("provider")
	ctx.Request = ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), "provider", provider))
	user, err := c.callbackfunc(ctx.Writer, ctx.Request)
	if err != nil {
		ctx.AbortWithError(http.StatusForbidden, gin.Error{Err: errors.New("something went wrong")})
		return
	}
	token, err := c.ProviderAuth(ctx, user)
	if err != nil {
		ctx.AbortWithError(http.StatusForbidden, gin.Error{Err: errors.New("something went wrong")})
		return
	}
	fragment := url.Values{}
	fragment.Set("token", token)
	fragmentString := fragment.Encode()
	redirectURL := c.redirectPage + "redirect#" + fragmentString
	ctx.Redirect(http.StatusFound, redirectURL)
}

// @Auth-with-provider			godoc
// @Summary						GET request for auth with provider
// @Description					requires param provider, for example google, facebook or apple  (at this moment apple not working) This request redirects to the provider's page for authorization, which in turn transmits a token in the parameters (token)
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
