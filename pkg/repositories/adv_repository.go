package repositories

import (
	"context"
	"fmt"

	"study_marketplace/database/queries"
	dbmappers "study_marketplace/pkg/domen/mappers/db_mappers"
	entities "study_marketplace/pkg/domen/models/entities"
	reqmodels "study_marketplace/pkg/domen/models/request_models"
)

type AdvertisementsRepository interface {
	CreateAdvertisement(ctx context.Context, arg *entities.Advertisement) (*entities.Advertisement, error)
	UpdateAdvertisement(ctx context.Context, arg *entities.Advertisement) (*entities.Advertisement, error)
	GetAdvertisementByID(ctx context.Context, id int64) (*entities.Advertisement, error)
	DeleteAdvertisementByID(ctx context.Context, advid, userid int64) error
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
	dbadv, err := t.q.CreateAdvertisement(ctx, dbmappers.AdvertisementToCreateAdvertisementParams(arg))
	if err != nil {
		return nil, err
	}
	return dbmappers.CreateAdvertisementRowToAdvertisement(dbadv), nil
}

func (t *advertisementsRepository) UpdateAdvertisement(ctx context.Context, arg *entities.Advertisement) (*entities.Advertisement, error) {
	dbadv, err := t.q.UpdateAdvertisement(ctx, dbmappers.AdvertisementToUpdateAdvertisementParams(arg))
	if err != nil {
		return nil, err
	}
	fullDbadv, err := t.q.GetAdvertisementCategoryAndUserByID(ctx, dbadv)
	if err != nil {
		return nil, err
	}
	return dbmappers.GetFullAdvToAdvertisement(fullDbadv), nil
}

func (t *advertisementsRepository) GetAdvertisementByID(ctx context.Context, id int64) (*entities.Advertisement, error) {
	fullDbadv, err := t.q.GetAdvertisementCategoryAndUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return dbmappers.GetFullAdvToAdvertisement(fullDbadv), nil
}

func (t *advertisementsRepository) DeleteAdvertisementByID(ctx context.Context, advid, userid int64) error {
	return t.q.DeleteAdvertisementByID(ctx, queries.DeleteAdvertisementByIDParams{
		ID:         advid,
		ProviderID: userid,
	})
}

func (t *advertisementsRepository) GetAdvertisementAll(ctx context.Context) ([]entities.Advertisement, error) {
	dbadvs, err := t.q.GetAdvertisementAll(ctx)
	if err != nil {
		return nil, err
	}
	return dbmappers.GetAdvertisementsAllToAdvertisements(dbadvs), nil
}

func (t *advertisementsRepository) FilterAdvertisements(ctx context.Context, filter *reqmodels.AdvertisementFilterRequest) (
	*entities.AdvertisementPagination, error) {
	arg := dbmappers.AdvertisementFiltRequestToFilterAdvertisementsParams(filter)
	fmt.Printf("arg: \n\n==========\n%+v\n", arg)
	filterdAdvs, err := t.q.FilterAdvertisements(ctx, arg)
	if err != nil {
		return nil, err
	}
	return dbmappers.FiltAdvToAdvPagination(&arg, filterdAdvs), nil
}

func (t *advertisementsRepository) GetAdvertisementMy(ctx context.Context, userID int64) ([]entities.Advertisement, error) {
	dbadv, err := t.q.GetMyAdvertisement(ctx, userID)
	if err != nil {
		return nil, err
	}
	return dbmappers.GetMyAdvertisementsToAdvertisements(dbadv), nil
}
