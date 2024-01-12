package services

import (
	"context"
	entities "study_marketplace/pkg/domain/models/entities"
	"study_marketplace/pkg/repositories"
)

type AuthService interface {
	ProviderAuth(ctx context.Context, userInfo *entities.User) (string, error)
}

type authService struct {
	db       repositories.AuthRepository
	genToken func(userid int64, userName string) (string, error)
}

func NewAuthService(genToken func(userid int64, userName string) (string, error), db repositories.AuthRepository) AuthService {
	return &authService{db, genToken}
}

func (s *authService) ProviderAuth(ctx context.Context, userInfo *entities.User) (string, error) {
	user, err := s.db.CreateorUpdateUser(ctx, userInfo)
	if err != nil {
		return "", err
	}
	token, err := s.genToken(user.ID, user.Email)
	if err != nil {
		return "", err
	}
	return token, nil
}
