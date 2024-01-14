package controllers

import (
	"net/http"

	reqm "study_marketplace/pkg/domain/mappers/reqresp_mappers"
	reqmodels "study_marketplace/pkg/domain/models/request_models"
	"study_marketplace/pkg/services"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	UserRegister(ctx *gin.Context)
	UserLogin(ctx *gin.Context)
	UserInfo(ctx *gin.Context)
	UserPatch(ctx *gin.Context)
	PasswordReset(ctx *gin.Context)
	PasswordCreate(ctx *gin.Context)
}

type userController struct {
	userService services.UserService
}

func NewUsersController(us services.UserService) UserController {
	return &userController{us}
}

// @Registraction	godoc
// @Summary			POST request for registration
// @Description		requires email and password for registration. Returns user info and Authorization token  in header
// @Tags			register
// @Accept			json
// @Produce			json
// @Param			user_info	body		reqmodels.RegistractionUserRequest	true	"user info for sign up"
// @Success			201			{object}	respmodels.SignUpINresponse
// @Failure			400			{object}	respmodels.FailedResponse
// @Router			/api/auth/register [post]
func (t *userController) UserRegister(ctx *gin.Context) {
	var inputModel reqmodels.RegistractionUserRequest
	if err := ctx.ShouldBindJSON(&inputModel); err != nil {
		ctx.JSON(http.StatusBadRequest, reqm.FailedResponse(err.Error()))
		return
	}
	if inputModel.Password == "" || inputModel.Email == "" {
		ctx.JSON(http.StatusBadRequest, reqm.FailedResponse("email and password required"))
		return
	}

	token, _, err := t.userService.UserRegister(ctx, reqm.RegUserToUser(&inputModel))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, reqm.FailedResponse(err.Error()))
		return
	}
	ctx.JSON(http.StatusCreated, reqm.TokenToSignUpINresponse(token))
}

// @Login			godoc
// @Summary			POST request for login
// @Description		requires email and password. Returns Authorization token in header as well
// @Tags			login
// @Accept			json
// @Produce			json
// @Param			request	body		reqmodels.LoginUserRequest	true	"request info"
// @Success			200		{object}	respmodels.SignUpINresponse
// @Failure			400		{object}	respmodels.FailedResponse
// @Router			/api/auth/login [post]
func (t *userController) UserLogin(ctx *gin.Context) {
	var inputModel reqmodels.LoginUserRequest
	if err := ctx.ShouldBindJSON(&inputModel); err != nil {
		ctx.JSON(http.StatusBadRequest, reqm.FailedResponse(err.Error()))
		return
	}
	if inputModel.Password == "" || inputModel.Email == "" {
		ctx.JSON(http.StatusBadRequest, reqm.FailedResponse("email and password required"))
		return
	}

	token, _, err := t.userService.UserLogin(ctx, reqm.LoginUserToUser(&inputModel))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, reqm.FailedResponse(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, reqm.TokenToSignUpINresponse(token))
}

// @Userinfo		godoc
// @Summary			Get request to see user info
// @Description		requires valid token
// @Tags			userinfo
// @Security		JWT
// @Param			Authorization	header	string	true	"Insert your access token"
// @Produce			json
// @Success			200	{object}	respmodels.UserInfoResponse
// @Failure			400	{object}	respmodels.FailedResponse
// @Router			/protected/userinfo [get]
func (t *userController) UserInfo(ctx *gin.Context) {
	userID := ctx.GetInt64("user_id")
	if userID == 0 {
		ctx.JSON(http.StatusBadRequest, reqm.FailedResponse("user id error"))
		return
	}
	user, err := t.userService.UserInfo(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, reqm.FailedResponse(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, reqm.UserToUserResponse(user))
}

// @User-patch		godoc
// @Summary			PATCH request to update user
// @Description		requires valid token and user info for update. Returns user info and Authorization token in header
// @Tags			user-patch
// @Security		JWT
// @Param			Authorization	header	string			true	"Insert your access token"
// @Param			userinfo		body	reqmodels.UpdateUserRequest		true	"user info for update"
// @Produce			json
// @Success			200	{object}	respmodels.SignUpINresponse
// @Failure			400	{object}	respmodels.FailedResponse
// @Router			/protected/user-patch [patch]
func (t *userController) UserPatch(ctx *gin.Context) {
	userId := ctx.GetInt64("user_id")
	var inputModel reqmodels.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&inputModel); err != nil {
		ctx.JSON(http.StatusBadRequest, reqm.FailedResponse(err.Error()))
		return
	}
	token, _, err := t.userService.UserPatch(ctx, reqm.UpdateUserRequestToUser(&inputModel, userId))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, reqm.FailedResponse(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, reqm.TokenToSignUpINresponse(token))
}

// @Reset-password	godoc
// @Summary			POST request to update password
// @Description		requires registered email address. TODO! This endpoint may not work
// @Tags			reset-password
// @Param			reset-password	body	reqmodels.PasswordResetRequest	true	"user email for update"
// @Produce			json
// @Success			200	{object}	respmodels.StringResponse
// @Failure			400	{object}	respmodels.FailedResponse
// @Router			/api/auth/reset-password [post]
func (t *userController) PasswordReset(ctx *gin.Context) {
	var userEmail reqmodels.PasswordResetRequest
	if err := ctx.ShouldBindJSON(&userEmail); err != nil {
		ctx.JSON(http.StatusBadRequest, reqm.FailedResponse(err.Error()))
		return
	}
	if userEmail.Email == "" {
		ctx.JSON(http.StatusBadRequest, reqm.FailedResponse("Email not provided."))
		return
	}
	if err := t.userService.PasswordReset(ctx, userEmail.Email); err != nil {
		ctx.JSON(http.StatusBadRequest, reqm.FailedResponse("Email not found."))
		return
	}
	ctx.JSON(http.StatusOK, reqm.StrResponse("Password Reset Email Has Been Sent"))
}

// @Create-password		godoc
// @Summary				PATCH request to create new password
// @Description			requires token. TODO! This endpoint may not work
// @Tags				create-password
// @Param				Authorization	header	string				true	"Insert your access token"
// @Param				create-password	body	reqmodels.PasswordCreateRequest	true	"new user password"
// @Produce				json
// @Success				200	{object}	respmodels.StringResponse
// @Failure				400	{object}	respmodels.FailedResponse
// @Router				/protected/create-password [patch]
func (t *userController) PasswordCreate(ctx *gin.Context) {
	userID := ctx.GetInt64("user_id")
	var newPassword reqmodels.PasswordCreateRequest
	if err := ctx.ShouldBindJSON(&newPassword); err != nil {
		ctx.JSON(http.StatusBadRequest, reqm.FailedResponse(err.Error()))
		return
	}
	if newPassword.Password == "" {
		ctx.JSON(http.StatusBadRequest, reqm.FailedResponse("New password not provided."))
		return
	}
	err := t.userService.PasswordCreate(ctx, userID, newPassword.Password)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, reqm.FailedResponse("Failed to create new password."))
		return
	}
	ctx.JSON(http.StatusOK, reqm.StrResponse("Password updated."))
}
