package repositories

import (
	"context"
	"fmt"

	"study_marketplace/database/queries"
	dbmappers "study_marketplace/pkg/domain/mappers/db_mappers"
	entities "study_marketplace/pkg/domain/models/entities"
)

//go:generate mockgen -destination=../../gen/mocks/mock_users_repository.go -package=mocks . UsersRepository

type UsersRepository interface {
	CreateUser(ctx context.Context, user *entities.User) (*entities.User, error)
	DeleteUser(ctx context.Context, id int64) error
	GetUserByEmail(ctx context.Context, email string) (*entities.User, error)
	GetUserById(ctx context.Context, id int64) (*entities.User, error)
	IsUserEmailExist(ctx context.Context, email string) (bool, error)
	ListUsers(ctx context.Context, arg queries.ListUsersParams) ([]queries.User, error)
	UpdateUser(ctx context.Context, user *entities.User) (*entities.User, error)
	CreateOrUpdateUser(ctx context.Context, user *entities.User) (*entities.User, error)
}

type usersRepository struct {
	q *queries.Queries
}

func NewUsersRepository(q *queries.Queries) *usersRepository { return &usersRepository{q} }

func (t *usersRepository) CreateUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	dbUser, err := t.q.CreateUser(ctx, dbmappers.UserToCreateUserParams(user))
	if err != nil {
		return nil, err
	}
	return dbmappers.QueryUserToUser(dbUser), nil
}

func (t *usersRepository) DeleteUser(ctx context.Context, id int64) error {
	return t.q.DeleteUser(ctx, id)
}

func (t *usersRepository) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	dbUser, err := t.q.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return dbmappers.QueryUserToUser(dbUser), nil
}

func (t *usersRepository) GetUserById(ctx context.Context, id int64) (*entities.User, error) {
	dbUser, err := t.q.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	return dbmappers.QueryUserToUser(dbUser), nil
}

func (t *usersRepository) IsUserEmailExist(ctx context.Context, email string) (bool, error) {
	return t.q.IsUserEmailExist(ctx, email)
}

func (t *usersRepository) ListUsers(ctx context.Context, arg queries.ListUsersParams) ([]queries.User, error) {
	return t.q.ListUsers(ctx, arg)
}

func (t *usersRepository) UpdateUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	dbUser, err := t.q.UpdateUser(ctx, dbmappers.UserToUpdateUserParams(user))
	if err != nil {
		return nil, err
	}
	return dbmappers.QueryUserToUser(dbUser), nil
}

func (t *usersRepository) CreateOrUpdateUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	dbUser, err := t.q.CreateOrUpdateUser(ctx, dbmappers.UserToCreateOrUpdateUser(user))
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, err
	}
	return dbmappers.CreateOrUpdateUserRowToUser(dbUser), nil
}
