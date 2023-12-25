package repositories

import (
	"context"

	"study_marketplace/database/queries"
	dbmappers "study_marketplace/domen/mappers/db_mappers"
	entities "study_marketplace/domen/models/entities_models"
)

type UsersRepository interface {
	CreateUser(ctx context.Context, arg queries.CreateUserParams) (*entities.User, error)
	DeleteUser(ctx context.Context, id int64) error
	GetUserByEmail(ctx context.Context, email string) (*entities.User, error)
	GetUserById(ctx context.Context, id int64) (*entities.User, error)
	IsUserEmailExist(ctx context.Context, email string) (bool, error)
	ListUsers(ctx context.Context, arg queries.ListUsersParams) ([]queries.User, error)
	UpdateUser(ctx context.Context, user *entities.User) (*entities.User, error)
	CreateorUpdateUser(ctx context.Context, user *entities.User) (*entities.User, error)
}

type usersRepository struct {
	q *queries.Queries
}

func NewUsersRepository(q *queries.Queries) *usersRepository { return &usersRepository{q} }

func (t *usersRepository) CreateUser(ctx context.Context, arg queries.CreateUserParams) (*entities.User, error) {
	dbuser, err := t.q.CreateUser(ctx, arg)
	if err != nil {
		return nil, err
	}
	return dbmappers.QueryUserToUser(dbuser), nil
}

func (t *usersRepository) DeleteUser(ctx context.Context, id int64) error {
	return t.q.DeleteUser(ctx, id)
}

func (t *usersRepository) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	dbuser, err := t.q.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return dbmappers.QueryUserToUser(dbuser), nil
}

func (t *usersRepository) GetUserById(ctx context.Context, id int64) (*entities.User, error) {
	dbuser, err := t.q.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	return dbmappers.QueryUserToUser(dbuser), nil
}

func (t *usersRepository) IsUserEmailExist(ctx context.Context, email string) (bool, error) {
	return t.q.IsUserEmailExist(ctx, email)
}

func (t *usersRepository) ListUsers(ctx context.Context, arg queries.ListUsersParams) ([]queries.User, error) {
	return t.q.ListUsers(ctx, arg)
}

func (t *usersRepository) UpdateUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	dbuser, err := t.q.UpdateUser(ctx, dbmappers.UserToUpdateUserParams(user))
	if err != nil {
		return nil, err
	}
	return dbmappers.QueryUserToUser(dbuser), nil
}

func (t *usersRepository) CreateorUpdateUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	dbuser, err := t.q.CreateorUpdateUser(ctx, dbmappers.UserToCreateOrUpdateUser(user))
	if err != nil {
		return nil, err
	}
	return dbmappers.QueryUserToUser(dbuser), nil
}
