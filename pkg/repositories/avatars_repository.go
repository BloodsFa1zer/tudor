package repositories

import (
	"context"
	"errors"
	"study_marketplace/database/queries"
	dbmappers "study_marketplace/pkg/domain/mappers/db_mappers"
	entities "study_marketplace/pkg/domain/models/entities"
)

//go:generate mockgen -destination=../../gen/mocks/mock_avatars_repository.go -package=mocks . AvatarsRepository

type AvatarsRepository interface {
	CreateAvatar(ctx context.Context, arg *entities.Avatar) (*entities.Avatar, error)
	UpdateAvatar(ctx context.Context, arg *entities.Avatar) (*entities.Avatar, error)
	DeleteAvatar(ctx context.Context, arg *entities.Avatar) (*entities.Avatar, error)
	GetAvatarByProviderID(ctx context.Context, arg *entities.Avatar) (*entities.Avatar, error)
	GetAvatarByID(ctx context.Context, arg *entities.Avatar) (*entities.Avatar, error)
}

type avatarRepository struct {
	q *queries.Queries
}

func NewAvatarsRepository(q *queries.Queries) *avatarRepository { return &avatarRepository{q} }

func (t *avatarRepository) CreateAvatar(ctx context.Context, arg *entities.Avatar) (*entities.Avatar, error) {
	dbavatar, err := t.q.CreateAvatar(ctx, dbmappers.AvatarToCreateAvatar(arg))
	if err != nil {
		return nil, err
	}
	return dbmappers.DBAvatarToAvatar(&dbavatar), nil
}

func (t *avatarRepository) UpdateAvatar(ctx context.Context, arg *entities.Avatar) (*entities.Avatar, error) {
	dbavatar, err := t.q.UpdateAvatarByProviderID(ctx, dbmappers.AvatarToUpdateAvatar(arg))
	if err != nil {
		return nil, err
	}
	return dbmappers.DBAvatarToAvatar(&dbavatar), nil
}

func (t *avatarRepository) DeleteAvatar(ctx context.Context, id int64) error {
	count, err := t.q.DeleteAvatarByProviderID(ctx, id)
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("avatar not found")
	}
	return nil
}

func (t *avatarRepository) GetAvatarByProviderID(ctx context.Context, providerID int64) (*entities.Avatar, error) {
	dbavatar, err := t.q.GetAvatarByProviderID(ctx, providerID)
	if err != nil {
		return nil, err
	}
	return dbmappers.DBAvatarToAvatar(&dbavatar), nil
}

func (t *avatarRepository) GetAvatarByID(ctx context.Context, id int64) (*entities.Avatar, error) {
	dbavatar, err := t.q.GetAvatarByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return dbmappers.DBAvatarToAvatar(&dbavatar), nil
}
