package dbmappers

import (
	"study_marketplace/database/queries"
	entities "study_marketplace/pkg/domen/models/entities"
	reqmodels "study_marketplace/pkg/domen/models/request_models"

	"github.com/jackc/pgx/v5/pgtype"
)

func UserToCreateOrUpdateUser(user *entities.User) queries.CreateOrUpdateUserParams {
	return queries.CreateOrUpdateUserParams{
		Name:     StrTopgText(user.Name),
		Email:    user.Email,
		Photo:    StrTopgText(user.Photo),
		Verified: user.Verified,
		Role:     user.Role,
	}
}

func QueryUserToUser(user queries.User) *entities.User {
	return &entities.User{
		ID:        user.ID,
		Name:      user.Name.String,
		Email:     user.Email,
		Photo:     user.Photo.String,
		Verified:  user.Verified,
		Password:  user.Password.String,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func UserToCreateUserParams(user *entities.User) queries.CreateUserParams {
	return queries.CreateUserParams{
		Name:     StrTopgText(user.Name),
		Email:    user.Email,
		Password: StrTopgText(user.Password),
		Photo:    StrTopgText(user.Photo),
		Verified: user.Verified,
		Role:     user.Role,
	}
}

func UserToUpdateUserParams(user *entities.User) queries.UpdateUserParams {
	return queries.UpdateUserParams{
		ID:       user.ID,
		Name:     StrTopgText(user.Name),
		Email:    user.Email,
		Photo:    StrTopgText(user.Photo),
		Verified: user.Verified,
		Password: StrTopgText(user.Password),
		Role:     user.Role,
	}
}

func CreateOrUpdateUserRowToUser(row queries.CreateOrUpdateUserRow) *entities.User {
	return &entities.User{
		ID:        row.ID,
		Name:      row.Name.String,
		Email:     row.Email,
		Photo:     row.Photo.String,
		Verified:  row.Verified,
		Password:  row.Password.String,
		Role:      row.Role,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}
}

func StrTopgText(str string) pgtype.Text {
	if str != "" {
		return pgtype.Text{String: str, Valid: true}
	}
	return pgtype.Text{String: "", Valid: false}
}

func IntTopgInt4(i int32) pgtype.Int4 {
	if i != 0 {
		return pgtype.Int4{Int32: int32(i), Valid: true}
	}
	return pgtype.Int4{Int32: 0, Valid: false}
}

func ParamListUsersToDbParam(param reqmodels.UsersListRequest) queries.ListUsersParams {
	return queries.ListUsersParams{
		Offset: int32(param.Offset),
		Limit:  int32(param.Limit),
	}
}
