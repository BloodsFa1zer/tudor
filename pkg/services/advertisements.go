package services

import (
	"context"

	"study_marketplace/database/queries"
	entities "study_marketplace/pkg/domen/models/entities"
	reqmodels "study_marketplace/pkg/domen/models/request_models"
	"study_marketplace/pkg/repositories"
)

type AdvertisementService interface {
	AdvCreate(ctx context.Context, adv *entities.Advertisement) (*entities.Advertisement, error)
	AdvPatch(ctx context.Context, adv *entities.Advertisement) (*entities.Advertisement, error)
	AdvDelete(ctx context.Context, advId int64, userId int64) error
	AdvGetAll(ctx context.Context) ([]entities.Advertisement, error)
	AdvGetByID(ctx context.Context, id int64) (queries.Advertisement, error)
	AdvGetFiltered(ctx context.Context, filter *reqmodels.AdvertisementFilterRequest) (*entities.AdvertisementPagination, error)
	AdvGetMy(ctx context.Context, userID int64) ([]queries.Advertisement, error)
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
	err := t.db.DeleteAdvertisementByID(ctx, advId, userId)
	if err != nil {
		return err
	}

	return nil
}

func (t *advertisementService) AdvGetAll(ctx context.Context) ([]entities.Advertisement, error) {
	advertisements, err := t.db.GetAdvertisementAll(ctx)

	if err != nil {
		return nil, err
	}

	return advertisements, nil
}

func (t *advertisementService) AdvGetByID(ctx context.Context, id int64) (queries.Advertisement, error) {
	// adverAndCat, err := t.db.GetAdvertisementAndCategoryByID(ctx, id)

	// if err != nil {
	// 	return queries.Advertisement{}, err
	// }
	// advertisement :=
	// 	queries.Advertisement{
	// 		ID:    adverAndCat.ID,
	// 		Title: adverAndCat.Title,
	// 		// Provider:   adverAndCat.Provider,
	// 		ProviderID: adverAndCat.ProviderID,
	// 		Attachment: adverAndCat.Attachment,
	// 		Experience: adverAndCat.Experience,
	// 		CategoryID: adverAndCat.CategoryID,
	// 		Time:       adverAndCat.Time,
	// 		Price:      adverAndCat.Price,
	// 		Format:     adverAndCat.Format,
	// 		Language:   adverAndCat.Language,
	// 		CreatedAt:  adverAndCat.CreatedAt,
	// 	}
	// // advertisement.Category = fmt.Sprintf("%s: %s", catFull.ParentName.String, catFull.CategoryName)
	// return advertisement, nil
	return queries.Advertisement{}, nil
}

func (t *advertisementService) AdvGetFiltered(ctx context.Context, filter *reqmodels.AdvertisementFilterRequest) (
	*entities.AdvertisementPagination, error) {
	advertisements, err := t.db.FilterAdvertisements(ctx, filter)
	if err != nil {
		return nil, err
	}
	return advertisements, nil
}

func (t *advertisementService) AdvGetMy(ctx context.Context, userID int64) ([]queries.Advertisement, error) {
	advertisements, err := t.db.GetAdvertisementMy(ctx, userID)

	if err != nil {
		return nil, err
	}

	return advertisements, nil
}
