package repositories

import (
	"context"

	"study_marketplace/internal/database/queries"
)

type AdvertisementsRepository interface {
	CreateAdvertisement(ctx context.Context, arg queries.CreateAdvertisementParams) (queries.Advertisement, error)
	UpdateAdvertisement(ctx context.Context, arg queries.UpdateAdvertisementParams) (queries.Advertisement, error)
	GetAdvertisementByID(ctx context.Context, id int64) (queries.Advertisement, error)
	DeleteAdvertisementByID(ctx context.Context, id int64) error
	GetAdvertisementAll(ctx context.Context) ([]queries.Advertisement, error)
	FilterAdvertisements(ctx context.Context, arg queries.FilterAdvertisementsParams) ([]queries.FilterAdvertisementsRow, error)
	GetAdvertisementMy(ctx context.Context, userID int64) ([]queries.Advertisement, error)
}

type advertisementsRepository struct {
	q *queries.Queries
}

func NewAdvertisementsRepository(q *queries.Queries) AdvertisementsRepository {
	return &advertisementsRepository{q}
}

func (t *advertisementsRepository) CreateAdvertisement(ctx context.Context, arg queries.CreateAdvertisementParams) (queries.Advertisement, error) {
	return t.q.CreateAdvertisement(ctx, arg)
}

func (t *advertisementsRepository) UpdateAdvertisement(ctx context.Context, arg queries.UpdateAdvertisementParams) (queries.Advertisement, error) {
	return t.q.UpdateAdvertisement(ctx, arg)
}

func (t *advertisementsRepository) GetAdvertisementByID(ctx context.Context, id int64) (queries.Advertisement, error) {
	return t.q.GetAdvertisementByID(ctx, id)
}

func (t *advertisementsRepository) DeleteAdvertisementByID(ctx context.Context, id int64) error {
	return t.q.DeleteAdvertisementByID(ctx, id)
}

func (t *advertisementsRepository) GetAdvertisementAll(ctx context.Context) ([]queries.Advertisement, error) {
	return t.q.GetAdvertisementAll(ctx)
}

func (t *advertisementsRepository) FilterAdvertisements(ctx context.Context, arg queries.FilterAdvertisementsParams) ([]queries.FilterAdvertisementsRow, error) {
	return t.q.FilterAdvertisements(ctx, arg)
}

func (t *advertisementsRepository) GetAdvertisementMy(ctx context.Context, userID int64) ([]queries.Advertisement, error) {
	return t.q.GetAdvertisementMy(ctx, userID)
}
