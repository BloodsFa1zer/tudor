package repositories

import (
	"context"
	"fmt"
	"study_marketplace/database/queries"
	dbmappers "study_marketplace/pkg/domain/mappers/db_mappers"
	entities "study_marketplace/pkg/domain/models/entities"
)

//go:generate mockgen -destination=../../gen/mocks/mock_auth_repository.go -package=mocks . AuthRepository

type AuthRepository interface {
	CreateOrUpdateUser(ctx context.Context, user *entities.User) (*entities.User, error)
}

type authRepository struct {
	q *queries.Queries
}

func NewAuthRepository(q *queries.Queries) *authRepository { return &authRepository{q} }

func (t *authRepository) CreateOrUpdateUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	dbUser, err := t.q.CreateOrUpdateUser(ctx, dbmappers.UserToCreateOrUpdateUser(user))
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, err
	}
	return dbmappers.CreateOrUpdateUserRowToUser(dbUser), nil
}
