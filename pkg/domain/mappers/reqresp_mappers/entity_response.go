package reqm

import (
	entities "study_marketplace/pkg/domain/models/entities"
	respmodels "study_marketplace/pkg/domain/models/response_models"

	"github.com/markbates/goth"
)

func GothToUserToUser(gothUser goth.User) *entities.User {
	return &entities.User{
		Name:     gothUser.Name,
		Email:    gothUser.Email,
		Photo:    gothUser.AvatarURL,
		Verified: true,
		Role:     "user",
	}
}

func FailedResponse(reason string) *respmodels.FailedResponse {
	return &respmodels.FailedResponse{
		Data:   reason,
		Status: "failed",
	}
}

func StrResponse(reason string) *respmodels.StringResponse {
	return &respmodels.StringResponse{
		Data:   reason,
		Status: "success",
	}
}
