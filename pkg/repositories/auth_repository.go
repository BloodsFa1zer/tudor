package repositories

import (
	"context"
	"fmt"
	"study_marketplace/database/queries"
	dbmappers "study_marketplace/pkg/domen/mappers/db_mappers"
	entities "study_marketplace/pkg/domen/models/entities"
)

type AuthRepository interface {
	CreateorUpdateUser(ctx context.Context, user *entities.User) (*entities.User, error)
}

type authRepository struct {
	q *queries.Queries
}

func NewAuthRepository(q *queries.Queries) *authRepository { return &authRepository{q} }

func (t *authRepository) CreateorUpdateUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	dbuser, err := t.q.CreateOrUpdateUser(ctx, dbmappers.UserToCreateOrUpdateUser(user))
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, err
	}
	return dbmappers.CreateOrUpdateUserRowToUser(dbuser), nil
}
