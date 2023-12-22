package tests_reposotiries

import (
	"context"

	"study_marketplace/database/queries"
)

type ReturnCreateUserModel struct {
	User  queries.User
	Error error
}

type ReturnGetUserByEmailModel struct {
	User  queries.User
	Error error
}

type GetUserByIdModel struct {
	User  queries.User
	Error error
}

type IsUserEmailExistModel struct {
	Exist bool
	Error error
}

type ListUsersModel struct {
	Users []queries.User
	Error error
}

type UpdateUserModel struct {
	User  queries.User
	Error error
}

type TestUsersRepository struct {
	ReturnCreateUser       ReturnCreateUserModel
	ReturnDeleteUser       error
	ReturnGetUserByEmail   ReturnGetUserByEmailModel
	ReturnGetUserById      GetUserByIdModel
	ReturnIsUserEmailExist IsUserEmailExistModel
	ReturnListUsers        ListUsersModel
	ReturnUpdateUser       UpdateUserModel
}

func (t *TestUsersRepository) CreateUser(ctx context.Context, arg queries.CreateUserParams) (queries.User, error) {
	return t.ReturnCreateUser.User, t.ReturnCreateUser.Error
}

func (t *TestUsersRepository) DeleteUser(ctx context.Context, id int64) error {
	return t.ReturnDeleteUser
}

func (t *TestUsersRepository) GetUserByEmail(ctx context.Context, email string) (queries.User, error) {
	return t.ReturnGetUserByEmail.User, t.ReturnGetUserByEmail.Error
}

func (t *TestUsersRepository) GetUserById(ctx context.Context, id int64) (queries.User, error) {
	return t.ReturnGetUserById.User, t.ReturnGetUserById.Error
}

func (t *TestUsersRepository) IsUserEmailExist(ctx context.Context, email string) (bool, error) {
	return t.ReturnIsUserEmailExist.Exist, t.ReturnIsUserEmailExist.Error
}

func (t *TestUsersRepository) ListUsers(ctx context.Context, arg queries.ListUsersParams) ([]queries.User, error) {
	return t.ReturnListUsers.Users, t.ReturnListUsers.Error
}

func (t *TestUsersRepository) UpdateUser(ctx context.Context, arg queries.UpdateUserParams) (queries.User, error) {
	return t.ReturnUpdateUser.User, t.ReturnUpdateUser.Error
}

func NewTestUsersRepository() *TestUsersRepository {
	return &TestUsersRepository{}
}
