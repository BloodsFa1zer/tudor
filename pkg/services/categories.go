package services

import (
	"context"
	"study_marketplace/database/queries"
	"study_marketplace/pkg/repositories"
)

type CategoriesService interface {
	CatGetAll(ctx context.Context) ([]queries.GetCategoriesWithChildrenRow, error)
	CatGetByID(ctx context.Context, id int) (queries.Category, error)
	CatGetByName(ctx context.Context, name string) (queries.Category, error)
	CatGetFullName(ctx context.Context, name string) (queries.GetCategoryAndParentRow, error)
	CatGetParets(ctx context.Context) ([]queries.Category, error)
}

type categoriesService struct {
	db repositories.CategoriesRepository
}

func NewCategoriesService(db repositories.CategoriesRepository) *categoriesService {
	return &categoriesService{db}
}

func (t *categoriesService) CatGetAll(ctx context.Context) ([]queries.GetCategoriesWithChildrenRow, error) {
	categories, err := t.db.GetCategoriesWithChildren(ctx)

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (t *categoriesService) CatGetByID(ctx context.Context, id int) (queries.Category, error) {
	category, err := t.db.GetCategoryByID(ctx, id)

	if err != nil {
		return queries.Category{}, err
	}

	return category, nil
}

func (t *categoriesService) CatGetByName(ctx context.Context, name string) (queries.Category, error) {
	category, err := t.db.GetCategoryByName(ctx, name)

	if err != nil {
		return queries.Category{}, err
	}

	return category, nil
}

func (t *categoriesService) CatGetFullName(ctx context.Context, name string) (queries.GetCategoryAndParentRow, error) {
	categoryName, err := t.db.GetCategoryAndParent(ctx, name)

	if err != nil {
		return queries.GetCategoryAndParentRow{}, err
	}

	return categoryName, nil
}

func (t *categoriesService) CatGetParets(ctx context.Context) ([]queries.Category, error) {
	parents, err := t.db.GetCategoryParents(ctx)

	if err != nil {
		return []queries.Category{}, err
	}

	return parents, nil
}
