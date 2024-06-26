package repositories

import (
	"context"
	"errors"

	"study_marketplace/database/queries"
	dbmappers "study_marketplace/pkg/domain/mappers/db_mappers"
	entities "study_marketplace/pkg/domain/models/entities"
	reqmodels "study_marketplace/pkg/domain/models/request_models"
)

//go:generate mockgen -destination=../../gen/mocks/mock_advertisements_repository.go -package=mocks . AdvertisementsRepository

type AdvertisementsRepository interface {
	CreateAdvertisement(ctx context.Context, arg *entities.Advertisement) (*entities.Advertisement, error)
	UpdateAdvertisement(ctx context.Context, arg *entities.Advertisement) (*entities.Advertisement, error)
	GetAdvertisementByID(ctx context.Context, id int64) (*entities.Advertisement, error)
	DeleteAdvertisementByID(ctx context.Context, advId, userId int64) error
	GetAdvertisementAll(ctx context.Context) ([]entities.Advertisement, error)
	FilterAdvertisements(ctx context.Context, filter *reqmodels.AdvertisementFilterRequest) (*entities.AdvertisementPagination, error)
	GetAdvertisementMy(ctx context.Context, userID int64) ([]entities.Advertisement, error)
}

type advertisementsRepository struct {
	q *queries.Queries
}

func NewAdvertisementsRepository(q *queries.Queries) AdvertisementsRepository {
	return &advertisementsRepository{q}
}

func (t *advertisementsRepository) CreateAdvertisement(ctx context.Context, arg *entities.Advertisement) (
	*entities.Advertisement, error) {
	dbAdv, err := t.q.CreateAdvertisement(ctx, dbmappers.AdvertisementToCreateAdvertisementParams(arg))
	if err != nil {
		return nil, err
	}
	return dbmappers.CreateAdvertisementRowToAdvertisement(dbAdv), nil
}

func (t *advertisementsRepository) UpdateAdvertisement(ctx context.Context, arg *entities.Advertisement) (*entities.Advertisement, error) {
	dbAdv, err := t.q.UpdateAdvertisement(ctx, dbmappers.AdvertisementToUpdateAdvertisementParams(arg))
	if err != nil {
		return nil, err
	}
	fullDbAdv, err := t.q.GetAdvertisementCategoryAndUserByID(ctx, dbAdv)
	if err != nil {
		return nil, err
	}
	return dbmappers.GetFullAdvToAdvertisement(fullDbAdv), nil
}

func (t *advertisementsRepository) GetAdvertisementByID(ctx context.Context, id int64) (*entities.Advertisement, error) {
	fullDbAdv, err := t.q.GetAdvertisementCategoryAndUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return dbmappers.GetFullAdvToAdvertisement(fullDbAdv), nil
}

func (t *advertisementsRepository) DeleteAdvertisementByID(ctx context.Context, advId, userid int64) error {
	count, err := t.q.DeleteAdvertisementByID(ctx, queries.DeleteAdvertisementByIDParams{
		ID:         advId,
		ProviderID: userid,
	})
	if err != nil {
		return err
	}

	if count == 0 {
		return errors.New("you have no rights to delete this adv or such adv doesn't exist")
	}
	return nil
}

func (t *advertisementsRepository) GetAdvertisementAll(ctx context.Context) ([]entities.Advertisement, error) {
	dbAdvs, err := t.q.GetAdvertisementAll(ctx)
	if err != nil {
		return nil, err
	}
	return dbmappers.GetAdvertisementsAllToAdvertisements(dbAdvs), nil
}

func (t *advertisementsRepository) FilterAdvertisements(ctx context.Context, filter *reqmodels.AdvertisementFilterRequest) (
	*entities.AdvertisementPagination, error) {
	arg := dbmappers.AdvertisementFilterRequestToFilterAdvertisementsParams(filter)
	filterAdvs, err := t.q.FilterAdvertisements(ctx, arg)
	if err != nil {
		return nil, err
	}
	return dbmappers.FilterAdvToAdvPagination(&arg, filterAdvs), nil
}

func (t *advertisementsRepository) GetAdvertisementMy(ctx context.Context, userID int64) ([]entities.Advertisement, error) {
	dbAdv, err := t.q.GetMyAdvertisement(ctx, userID)
	if err != nil {
		return nil, err
	}
	return dbmappers.GetMyAdvertisementsToAdvertisements(dbAdv), nil
}
