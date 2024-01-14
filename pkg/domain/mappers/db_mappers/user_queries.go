package dbmappers

import (
	"study_marketplace/database/queries"
	entities "study_marketplace/pkg/domain/models/entities"
	reqmodels "study_marketplace/pkg/domain/models/request_models"

	"github.com/jackc/pgx/v5/pgtype"
)

func UserToCreateOrUpdateUser(user *entities.User) queries.CreateOrUpdateUserParams {
	return queries.CreateOrUpdateUserParams{
		Name:       StrTopgText(user.Name),
		Email:      user.Email,
		Photo:      StrTopgText(user.Photo),
		Verified:   user.Verified,
		Password:   StrTopgText(user.Password),
		Role:       user.Role,
		Name_2:     StrTopgText(user.Name),
		Photo_2:    StrTopgText(user.Photo),
		Verified_2: BoolTopgBool(user.Verified),
		Password_2: StrTopgText(user.Password),
		Role_2:     StrTopgText(user.Role),
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
		ID:      user.ID,
		Column2: user.Name,
		Column3: user.Email,
		Column4: user.Password,
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

func BoolTopgBool(b bool) pgtype.Bool {
	if b {
		return pgtype.Bool{Bool: b, Valid: true}
	}
	return pgtype.Bool{Bool: false, Valid: false}
}

func ParamListUsersToDbParam(param reqmodels.UsersListRequest) queries.ListUsersParams {
	return queries.ListUsersParams{
		Offset: int32(param.Offset),
		Limit:  int32(param.Limit),
	}
}
