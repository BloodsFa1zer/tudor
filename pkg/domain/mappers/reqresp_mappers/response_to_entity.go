package reqm

import (
	entities "study_marketplace/pkg/domain/models/entities"

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
