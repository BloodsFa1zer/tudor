package controllers

import (
	"context"
	"errors"
	"net/http"
	reqm "study_marketplace/pkg/domen/mappers/reqresp_mappers"
	respmodels "study_marketplace/pkg/domen/models/response_models"
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
	user, err := gothic.CompleteUserAuth(ctx.Writer, ctx.Request)
	if err != nil {
		ctx.AbortWithError(http.StatusForbidden, gin.Error{Err: errors.New("something went wrong")})
		return
	}
	token, err := c.ProviderAuth(ctx, reqm.GothToUserToUser(user))
	if err != nil {
		ctx.AbortWithError(http.StatusForbidden, gin.Error{Err: errors.New("something went wrong")})
		return
	}
	ctx.Header("Authorization", token)
	ctx.JSON(http.StatusOK, respmodels.StringResponse{Data: token, Status: "success"})
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