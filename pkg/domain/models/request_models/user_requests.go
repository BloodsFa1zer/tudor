package reqmodels

// RegistractionUserRequest godoc
type RegistractionUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginUserRequest godoc
type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UpdateUserRequest godoc
type UpdateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// PasswordResetRequest godoc
type PasswordResetRequest struct {
	Email string `json:"email"`
}

// PasswordChangeRequest godoc
type PasswordChangeRequest struct {
	CurrentPassword string `json:"currentPassword"`
	NewPassword     string `json:"newPassword"`
}

// PasswordCreateRequest godoc
type PasswordCreateRequest struct {
	Password string `json:"password"`
}

type UsersListRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}
