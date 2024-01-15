package reqm

import (
	entities "study_marketplace/pkg/domain/models/entities"
	reqmodels "study_marketplace/pkg/domain/models/request_models"
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

func UpdateUserRequestToUser(updateUser *reqmodels.UpdateUserRequest, id int64) *entities.User {
	return &entities.User{
		ID:       id,
		Name:     updateUser.Name,
		Email:    updateUser.Email,
		Verified: true,
		Role:     "user",
	}
}
