package reqm

import (
	entities "study_marketplace/pkg/domen/models/entities"
	respmodels "study_marketplace/pkg/domen/models/response_models"
)

func UserToUserResponse(user *entities.User) *respmodels.UserInfoResponse {
	return &respmodels.UserInfoResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Photo:     user.Photo,
		Verified:  user.Verified,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
