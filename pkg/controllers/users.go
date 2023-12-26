package controllers

import (
	"net/http"

	reqm "study_marketplace/pkg/domen/mappers/reqresp_mappers"
	reqmodels "study_marketplace/pkg/domen/models/request_models"
	respmodels "study_marketplace/pkg/domen/models/response_models"
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
// @Description		requires email and password for registration. Returns user info and in header Authorization token
// @Tags			register
// @Accept			json
// @Produce			json
// @Param			user_info	body		reqmodels.RegistractionUserRequest	true	"user info for sign in"
// @Success			200			{object}	map[string]interface{}
// @Router			/api/auth/register [post]
func (t *userController) UserRegister(ctx *gin.Context) {
	var inputModel reqmodels.RegistractionUserRequest
	if err := ctx.ShouldBindJSON(&inputModel); err != nil {
		ctx.JSON(http.StatusBadRequest, respmodels.NewResponseFailed(err.Error()))
		return
	}
	if inputModel.Password == "" || inputModel.Email == "" {
		ctx.JSON(http.StatusBadRequest, respmodels.NewResponseFailed("email and password required"))
		return
	}

	token, user, err := t.userService.UserRegister(ctx, reqm.RegUserToUser(&inputModel))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, respmodels.NewResponseFailed(err.Error()))
		return
	}
	ctx.Request.Header.Add("Authorization", token)
	ctx.JSON(http.StatusCreated, respmodels.NewResponseSuccess(user))
}

// @Login			godoc
// @Summary			POST request for login
// @Description		requires email and password.  Returns token and in header Authorization token as well
// @Tags			login
// @Accept			json
// @Produce			json
// @Param			request	body		reqmodels.LoginUserRequest	true	"request info"
// @Success			200		{object}	map[string]interface{}
// @Router			/api/auth/login [post]
func (t *userController) UserLogin(ctx *gin.Context) {
	var inputModel reqmodels.LoginUserRequest
	if err := ctx.ShouldBindJSON(&inputModel); err != nil {
		ctx.JSON(http.StatusBadRequest, respmodels.NewResponseFailed(err.Error()))
		return
	}
	if inputModel.Password == "" || inputModel.Email == "" {
		ctx.JSON(http.StatusBadRequest, respmodels.NewResponseFailed("email and password required"))
		return
	}

	token, err := t.userService.UserLogin(ctx, reqm.LoginUserToUser(&inputModel))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, respmodels.NewResponseFailed(err.Error()))
		return
	}
	ctx.Request.Header.Add("Authorization", token)
	ctx.JSON(http.StatusOK, respmodels.NewResponseSuccess(token))
}

// @Userinfo		godoc
// @Summary			Get request to see user info
// @Description		requires valid token
// @Tags			userinfo
// @Security		JWT
// @Param			Authorization	header	string	true	"Insert your access token"
// @Produce			json
// @Success			200	{object}	map[string]interface{}
// @Router			/protected/userinfo [get]
func (t *userController) UserInfo(ctx *gin.Context) {
	userID := ctx.GetInt64("user_id")
	if userID == 0 {
		ctx.JSON(http.StatusBadRequest, respmodels.NewResponseFailed("user id error"))
		return
	}
	user, err := t.userService.UserInfo(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, respmodels.NewResponseFailed(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, respmodels.NewResponseSuccess(user))
}

// @User-patch		godoc
// @Summary			PATCH request to update user
// @Description		requires valid token
// @Tags			user-patch
// @Security		JWT
// @Param			Authorization	header	string			true	"Insert your access token"
// @Param			userinfo		body	reqmodels.UpdateUserRequest		true	"user info for update"
// @Produce			json
// @Success			200	{object}	map[string]interface{}
// @Router			/protected/user-patch [patch]
func (t *userController) UserPatch(ctx *gin.Context) {
	userId := ctx.GetInt64("user_id")
	var inputModel reqmodels.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&inputModel); err != nil {
		ctx.JSON(http.StatusBadRequest, respmodels.NewResponseFailed(err.Error()))
		return
	}
	inputModel.ID = userId
	user, err := t.userService.UserPatch(ctx, reqm.UpdateUserRequestToUser(&inputModel))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, respmodels.NewResponseFailed(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, respmodels.NewResponseSuccess(user))
}

// @Reset-password	godoc
// @Summary			POST request to update password
// @Description		requires registred email address
// @Tags			reset-password
// @Param			reset-password	body	reqmodels.PasswordResetRequest	true	"user email for update"
// @Produce			json
// @Success			200	{object}	map[string]interface{}
// @Router			/api/auth/reset-password [post]
func (t *userController) PasswordReset(ctx *gin.Context) {
	var userEmail reqmodels.PasswordResetRequest
	if err := ctx.ShouldBindJSON(&userEmail); err != nil {
		ctx.JSON(http.StatusBadRequest, respmodels.NewResponseFailed("Can't read email."))
		return
	}
	if userEmail.Email == "" {
		ctx.JSON(http.StatusBadRequest, respmodels.NewResponseFailed("Email not provided."))
		return
	}
	_, err := t.userService.PasswordReset(ctx, userEmail.Email)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, respmodels.NewResponseFailed("Email not found."))
		return
	}

	ctx.JSON(http.StatusOK, respmodels.NewResponseSuccess("Password Reset Email Has Been Sent"))
}

// @Create-password		godoc
// @Summary				PATCH request to create new password
// @Description			requires token
// @Tags				create-password
// @Param				Authorization	header	string				true	"Insert your access token"
// @Param				create-password	body	reqmodels.PasswordCreateRequest	true	"new user password"
// @Produce				json
// @Success				200	{object}	map[string]interface{}
// @Router				/protected/create-password [patch]
func (t *userController) PasswordCreate(ctx *gin.Context) {
	userID := ctx.GetInt64("user_id")
	var newPassword reqmodels.PasswordCreateRequest
	if err := ctx.ShouldBindJSON(&newPassword); err != nil {
		ctx.JSON(http.StatusBadRequest, respmodels.NewResponseFailed("New password not provided."))
	}
	if newPassword.Password == "" {
		ctx.JSON(http.StatusBadRequest, respmodels.NewResponseFailed("New password not provided."))
		return
	}
	err := t.userService.PasswordCreate(ctx, userID, newPassword.Password)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, respmodels.NewResponseFailed("Failed to create new passowrd."))
		return
	}
	ctx.JSON(http.StatusOK, respmodels.NewResponseSuccess("Password updated."))
}

// // method used for password-middleware
// // won't be publick endpoint
// func (t *userController) GetPassword(ctx *gin.Context) string {
// 	userID := ctx.GetInt64("user_id")
// 	_, err := t.userService.UserInfo(ctx, userID)

// 	if err != nil {
// 		ctx.JSON(http.StatusUnauthorized, respmodels.NewResponseFailed("No user found."))
// 		return ""
// 	}

// 	return "user.Password"
// }
