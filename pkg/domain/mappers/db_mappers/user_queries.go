package dbmappers

import (
	"database/sql"
	"study_marketplace/database/queries"
	entities "study_marketplace/pkg/domain/models/entities"
	reqmodels "study_marketplace/pkg/domain/models/request_models"
)

func UserToCreateOrUpdateUser(user *entities.User) queries.CreateOrUpdateUserParams {
	return queries.CreateOrUpdateUserParams{
		Name:       StrToSqlStr(user.Name),
		Email:      user.Email,
		Photo:      StrToSqlStr(user.Photo),
		Verified:   user.Verified,
		Password:   StrToSqlStr(user.Password),
		Role:       user.Role,
		Name_2:     StrToSqlStr(user.Name),
		Photo_2:    StrToSqlStr(user.Photo),
		Verified_2: BoolTopgBool(user.Verified),
		Password_2: StrToSqlStr(user.Password),
		Role_2:     StrToSqlStr(user.Role),
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
		Name:     StrToSqlStr(user.Name),
		Email:    user.Email,
		Password: StrToSqlStr(user.Password),
		Photo:    StrToSqlStr(user.Photo),
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

func StrToSqlStr(str string) sql.NullString {
	if str != "" {
		return sql.NullString{String: str, Valid: true}
	}
	return sql.NullString{String: "", Valid: false}
}

func IntTopgInt4(i int32) sql.NullInt32 {
	if i != 0 {
		return sql.NullInt32{Int32: int32(i), Valid: true}
	}
	return sql.NullInt32{Int32: 0, Valid: false}
}

func BoolTopgBool(b bool) sql.NullBool {
	if b {
		return sql.NullBool{Bool: b, Valid: true}
	}
	return sql.NullBool{Bool: false, Valid: false}
}

func ParamListUsersToDbParam(param reqmodels.UsersListRequest) queries.ListUsersParams {
	return queries.ListUsersParams{
		Offset: int32(param.Offset),
		Limit:  int32(param.Limit),
	}
}
