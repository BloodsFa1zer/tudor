package reqm

import (
	entities "study_marketplace/pkg/domain/models/entities"
	respmodels "study_marketplace/pkg/domain/models/response_models"
)

func UserToUserResponse(user *entities.User) *respmodels.UserInfoResponse {
	return &respmodels.UserInfoResponse{
		ResponseUser: respmodels.ResponseUser{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Photo:     user.Photo,
			Verified:  user.Verified,
			Role:      user.Role,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
		Status: "success",
	}
}

func TokenToSignUpINresponse(token string) *respmodels.SignUpINresponse {
	return &respmodels.SignUpINresponse{
		AccessToken: struct {
			Token string `json:"token"`
		}{
			Token: token,
		},
		Status: "success",
	}
}
