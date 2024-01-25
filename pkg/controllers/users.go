package controllers

import (
	"fmt"
	"net/http"
	"os"

	reqm "study_marketplace/pkg/domain/mappers/reqresp_mappers"
	"study_marketplace/pkg/domain/models/entities"
	reqmodels "study_marketplace/pkg/domain/models/request_models"
	v "study_marketplace/pkg/infrastructure/validator"
	"study_marketplace/pkg/services"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	UserRegister(ctx *gin.Context)
	UserLogin(ctx *gin.Context)
	UserInfo(ctx *gin.Context)
	UserPatch(ctx *gin.Context)
	PasswordReset(ctx *gin.Context)
	PasswordChange(ctx *gin.Context)
	PasswordCreate(ctx *gin.Context)
	EmailChange(ctx *gin.Context)
	UploadAvatar(ctx *gin.Context)
}

type userController struct {
	userService services.UserService
	basicAppUrl string
}

func NewUsersController(us services.UserService, basicUrl string) UserController {
	return &userController{us, basicUrl}
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
	if err := v.Validate(inputModel); err != nil {
		ctx.JSON(http.StatusBadRequest, reqm.FailedResponse(err.Error()))
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
	if err := v.Validate(inputModel); err != nil {
		ctx.JSON(http.StatusBadRequest, reqm.FailedResponse(err.Error()))
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
	if err := v.Validate(inputModel); err != nil {
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
	if err := v.Validate(userEmail); err != nil {
		ctx.JSON(http.StatusBadRequest, reqm.FailedResponse(err.Error()))
		return
	}
	if err := t.userService.PasswordReset(ctx, userEmail.Email); err != nil {
		ctx.JSON(http.StatusBadRequest, reqm.FailedResponse("Email not found."))
		return
	}
	ctx.JSON(http.StatusOK, reqm.StrResponse("Password Reset Email Has Been Sent"))
}

// @Change-password	godoc
// @Summary			POST request to update password
// @Description		requires current password and new password
// @Tags			change-password
// @Param			Authorization	header	string				true	"Insert your access token"
// @Param			change-password	body	reqmodels.PasswordChangeRequest	true	"user email for update"
// @Produce			json
// @Success			200	{object}	respmodels.StringResponse
// @Failure			400	{object}	respmodels.FailedResponse
// @Router			/protected/change-password [post]
func (t *userController) PasswordChange(ctx *gin.Context) {
	var request reqmodels.PasswordChangeRequest
	userId := ctx.GetInt64("user_id")
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, reqm.FailedResponse("Unable to read the request."))
		return
	}

	if request.CurrentPassword == request.NewPassword {
		ctx.JSON(http.StatusBadRequest, reqm.FailedResponse("current password and new password are equal"))
		return
	}

	if err := v.Validate(request); err != nil {
		ctx.JSON(http.StatusBadRequest, reqm.FailedResponse(err.Error()))
		return
	}

	if err := t.userService.PasswordChange(ctx, userId, request.CurrentPassword, request.NewPassword); err != nil {
		ctx.JSON(http.StatusUnauthorized, reqm.FailedResponse(fmt.Sprintf("Password change failed: %s", err.Error())))
		return
	}
	ctx.JSON(http.StatusOK, reqm.StrResponse("Password has been updated"))
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
	if err := v.Validate(newPassword); err != nil {
		ctx.JSON(http.StatusBadRequest, reqm.FailedResponse(err.Error()))
		return
	}

	err := t.userService.PasswordCreate(ctx, userID, newPassword.Password)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, reqm.FailedResponse(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, reqm.StrResponse("Password updated."))
}

// @Change-email	godoc
// @Summary			POST request to update email
// @Description		requires current password and new email
// @Tags			change-email
// @Param			Authorization	header	string				true	"Insert your access token"
// @Param			change-email	body	reqmodels.EmailChangeRequest	true	"user email for update"
// @Produce			json
// @Success			200	{object}	respmodels.StringResponse
// @Failure			400	{object}	respmodels.FailedResponse
// @Router			/protected/change-email [post]
func (t *userController) EmailChange(ctx *gin.Context) {
	var request reqmodels.EmailChangeRequest
	userId := ctx.GetInt64("user_id")
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, reqm.FailedResponse("Unable to read the request."))
		return
	}

	if err := v.Validate(request); err != nil {
		ctx.JSON(http.StatusBadRequest, reqm.FailedResponse(err.Error()))
		return
	}

	if err := t.userService.EmailChange(ctx, userId, request.CurrentPassword, request.NewEmail); err != nil {
		ctx.JSON(http.StatusUnauthorized, reqm.FailedResponse(fmt.Sprintf("Email change failed: %s", err.Error())))
		return
	}
	ctx.JSON(http.StatusOK, reqm.StrResponse("Email has been updated"))
}

// @Upload-avatar	godoc
// @Summary			POST request to upload avatar
// @Description		requires valid token and avatar
// @Tags			upload-avatar
// @Security		JWT
// @Param			Authorization	header	string				true	"Insert your access token"
// @Param			avatar			formData	file				true	"avatar for upload"
// @Produce			json
// @Success			200	{object}	respmodels.StringResponse
// @Failure			400	{object}	respmodels.FailedResponse
// @Router			/protected/upload-avatar [post]
func (t *userController) UploadAvatar(ctx *gin.Context) {
	userID := ctx.GetInt64("user_id")
	file, err := ctx.FormFile("avatar")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, reqm.FailedResponse(err.Error()))
		return
	}
	user, err := t.userService.UserInfo(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, reqm.FailedResponse(err.Error()))
		return
	}
	if user.Photo != "" {
		_ = os.Remove(user.Photo)
	}

	path := fmt.Sprintf("avatars/%d-%s", userID, file.Filename)
	if err := ctx.SaveUploadedFile(file, path); err != nil {
		ctx.JSON(http.StatusBadRequest, reqm.FailedResponse(err.Error()))
		return
	}

	if _, _, err = t.userService.UserPatch(ctx,
		&entities.User{ID: userID, Photo: fmt.Sprintf("%s/%s", t.basicAppUrl, path)}); err != nil {
		ctx.JSON(http.StatusBadRequest, reqm.FailedResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, reqm.StrResponse("Avatar uploaded"))
}
