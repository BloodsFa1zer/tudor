package services

import (
	"context"
	entities "study_marketplace/pkg/domain/models/entities"
	"study_marketplace/pkg/repositories"
)

type AvatarsService interface {
	CreateAvatar(ctx context.Context, arg *entities.Avatar) (*entities.Avatar, error)
	UpdateAvatar(ctx context.Context, arg *entities.Avatar) (*entities.Avatar, error)
	DeleteAvatar(ctx context.Context, arg *entities.Avatar) (*entities.Avatar, error)
	GetAvatarByProviderID(ctx context.Context, arg *entities.Avatar) (*entities.Avatar, error)
	GetAvatarByID(ctx context.Context, arg *entities.Avatar) (*entities.Avatar, error)
}

type avatarsService struct {
	avatarsRepo repositories.AvatarsRepository
}

func NewAvatarsService(avatarsRepo repositories.AvatarsRepository) *avatarsService {
	return &avatarsService{avatarsRepo}
}

func (t *avatarsService) CreateAvatar(ctx context.Context, arg *entities.Avatar) (*entities.Avatar, error) {
	return t.avatarsRepo.CreateAvatar(ctx, arg)
}

func (t *avatarsService) UpdateAvatar(ctx context.Context, arg *entities.Avatar) (*entities.Avatar, error) {
	return t.avatarsRepo.UpdateAvatar(ctx, arg)
}

func (t *avatarsService) DeleteAvatar(ctx context.Context, arg *entities.Avatar) (*entities.Avatar, error) {
	return t.avatarsRepo.DeleteAvatar(ctx, arg)
}

func (t *avatarsService) GetAvatarByProviderID(ctx context.Context, arg *entities.Avatar) (*entities.Avatar, error) {
	return t.avatarsRepo.GetAvatarByProviderID(ctx, arg)
}

func (t *avatarsService) GetAvatarByID(ctx context.Context, arg *entities.Avatar) (*entities.Avatar, error) {
	return t.avatarsRepo.GetAvatarByID(ctx, arg)
}
