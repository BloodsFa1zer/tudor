package reqresmappers

import (
	entities "study_marketplace/domen/models/entities_models"
	reqmodels "study_marketplace/domen/models/request_models"
)

func RegUserToUser(regUser *reqmodels.RegistractionUserRequest) *entities.User {
	return &entities.User{
		Name:     regUser.Name,
		Email:    regUser.Email,
		Password: regUser.Password,
		Verified: true,
		Role:     "user",
	}
}

func LoginUserToUser(loginUser *reqmodels.LoginUserRequest) *entities.User {
	return &entities.User{
		Email:    loginUser.Email,
		Password: loginUser.Password,
	}
}

func UpdateUserRequestToUser(updateUser *reqmodels.UpdateUserRequest) *entities.User {
	return &entities.User{
		ID:    updateUser.ID,
		Name:  updateUser.Name,
		Email: updateUser.Email,
	}
}
