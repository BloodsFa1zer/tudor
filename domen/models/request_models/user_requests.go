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
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// PasswordResetRequest godoc
type PasswordResetRequest struct {
	Email string `json:"email"`
}

// PasswordCreateRequest godoc
type PasswordCreateRequest struct {
	Password string `json:"password"`
}
