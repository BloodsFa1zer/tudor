package repositories

import (
	"context"

	"study_marketplace/database/queries"
)

type UsersRepository interface {
	CreateUser(ctx context.Context, arg queries.CreateUserParams) (queries.User, error)
	DeleteUser(ctx context.Context, id int64) error
	GetUserByEmail(ctx context.Context, email string) (queries.User, error)
	GetUserById(ctx context.Context, id int64) (queries.User, error)
	IsUserEmailExist(ctx context.Context, email string) (bool, error)
	ListUsers(ctx context.Context, arg queries.ListUsersParams) ([]queries.User, error)
	UpdateUser(ctx context.Context, arg queries.UpdateUserParams) (queries.User, error)
}

type usersRepository struct {
	q *queries.Queries
}

func NewUsersRepository(q *queries.Queries) *usersRepository { return &usersRepository{q} }

func (t *usersRepository) CreateUser(ctx context.Context, arg queries.CreateUserParams) (queries.User, error) {
	return t.q.CreateUser(ctx, arg)
}

func (t *usersRepository) DeleteUser(ctx context.Context, id int64) error {
	return t.q.DeleteUser(ctx, id)
}

func (t *usersRepository) GetUserByEmail(ctx context.Context, email string) (queries.User, error) {
	return t.q.GetUserByEmail(ctx, email)
}

func (t *usersRepository) GetUserById(ctx context.Context, id int64) (queries.User, error) {
	return t.q.GetUserById(ctx, id)
}

func (t *usersRepository) IsUserEmailExist(ctx context.Context, email string) (bool, error) {
	return t.q.IsUserEmailExist(ctx, email)
}

func (t *usersRepository) ListUsers(ctx context.Context, arg queries.ListUsersParams) ([]queries.User, error) {
	return t.q.ListUsers(ctx, arg)
}

func (t *usersRepository) UpdateUser(ctx context.Context, arg queries.UpdateUserParams) (queries.User, error) {
	return t.q.UpdateUser(ctx, arg)
}
