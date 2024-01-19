package services

import (
	"context"

	entities "study_marketplace/pkg/domain/models/entities"
	reqmodels "study_marketplace/pkg/domain/models/request_models"
	"study_marketplace/pkg/repositories"
)

type AdvertisementService interface {
	AdvCreate(ctx context.Context, adv *entities.Advertisement) (*entities.Advertisement, error)
	AdvPatch(ctx context.Context, adv *entities.Advertisement) (*entities.Advertisement, error)
	AdvDelete(ctx context.Context, advId int64, userId int64) error
	AdvGetAll(ctx context.Context) ([]entities.Advertisement, error)
	AdvGetByID(ctx context.Context, id int64) (*entities.Advertisement, error)
	AdvGetFiltered(ctx context.Context, filter *reqmodels.AdvertisementFilterRequest) (*entities.AdvertisementPagination, error)
	AdvGetMy(ctx context.Context, userID int64) ([]entities.Advertisement, error)
}

type advertisementService struct {
	db repositories.AdvertisementsRepository
}

func NewAdvertisementService(db repositories.AdvertisementsRepository) AdvertisementService {
	return &advertisementService{db}
}

func (t *advertisementService) AdvCreate(ctx context.Context, adv *entities.Advertisement) (*entities.Advertisement, error) {
	return t.db.CreateAdvertisement(ctx, adv)
}

func (t *advertisementService) AdvPatch(ctx context.Context, adv *entities.Advertisement) (*entities.Advertisement, error) {

	return t.db.UpdateAdvertisement(ctx, adv)
}

func (t *advertisementService) AdvDelete(ctx context.Context, advId int64, userId int64) error {
	return t.db.DeleteAdvertisementByID(ctx, advId, userId)
}

func (t *advertisementService) AdvGetAll(ctx context.Context) ([]entities.Advertisement, error) {
	return t.db.GetAdvertisementAll(ctx)
}

func (t *advertisementService) AdvGetByID(ctx context.Context, id int64) (*entities.Advertisement, error) {
	return t.db.GetAdvertisementByID(ctx, id)
}

func (t *advertisementService) AdvGetFiltered(ctx context.Context, filter *reqmodels.AdvertisementFilterRequest) (
	*entities.AdvertisementPagination, error) {
	return t.db.FilterAdvertisements(ctx, filter)
}

func (t *advertisementService) AdvGetMy(ctx context.Context, userID int64) ([]entities.Advertisement, error) {
	return t.db.GetAdvertisementMy(ctx, userID)
}
