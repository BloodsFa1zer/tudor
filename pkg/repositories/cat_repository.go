package repositories

import (
	"context"

	"study_marketplace/database/queries"
)

type CategoriesRepository interface {
	GetCategoriesWithChildren(ctx context.Context) ([]queries.GetCategoriesWithChildrenRow, error)
	GetCategoryAndParent(ctx context.Context, name string) (queries.GetCategoryAndParentRow, error)
	GetCategoryByID(ctx context.Context, id int32) (queries.Category, error)
	GetCategoryByName(ctx context.Context, name string) (queries.Category, error)
	GetCategoryParents(ctx context.Context) ([]queries.Category, error)
}

type categoriesRepository struct {
	q *queries.Queries
}

func NewCategoriesRepository(q *queries.Queries) *categoriesRepository {
	return &categoriesRepository{q}
}

func (t *categoriesRepository) GetCategoriesWithChildren(ctx context.Context) ([]queries.GetCategoriesWithChildrenRow, error) {
	return t.q.GetCategoriesWithChildren(ctx)
}

func (t *categoriesRepository) GetCategoryAndParent(ctx context.Context, name string) (queries.GetCategoryAndParentRow, error) {
	return t.q.GetCategoryAndParent(ctx, name)
}

func (t *categoriesRepository) GetCategoryByID(ctx context.Context, id int32) (queries.Category, error) {
	return t.q.GetCategoryByID(ctx, id)
}

func (t *categoriesRepository) GetCategoryByName(ctx context.Context, name string) (queries.Category, error) {
	return t.q.GetCategoryByName(ctx, name)
}

func (t *categoriesRepository) GetCategoryParents(ctx context.Context) ([]queries.Category, error) {
	return t.q.GetCategoryParents(ctx)
}
