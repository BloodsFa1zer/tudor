package reqmodels

// RegistrationUserRequest godoc
type RegistrationUserRequest struct {
	Name     string `json:"name" validate:"min=2,max=50, nonzero"`
	Email    string `json:"email" validate:"email,nonzero"`
	Password string `json:"password" validate:"password"`
}

// LoginUserRequest godoc
type LoginUserRequest struct {
	Email    string `json:"email" validate:"nonzero"`
	Password string `json:"password" validate:"nonzero"`
}

// UpdateUserRequest godoc
type UpdateUserRequest struct {
	Name  string `json:"name" validate:"min=2,max=50, nonzero"` // deprecated
	Email string `json:"email" validate:"email,nonzero"`        // deprecated
}

// PasswordResetRequest godoc
type PasswordResetRequest struct {
	Email string `json:"email" validate:"email,nonzero"`
}

// PasswordChangeRequest godoc
type PasswordChangeRequest struct {
	CurrentPassword string `json:"currentPassword" validate:"password"`
	NewPassword     string `json:"newPassword" validate:"password"`
}

// EmailChangeRequest godoc
type EmailChangeRequest struct {
	CurrentPassword string `json:"currentPassword" validate:"password"`
	NewEmail        string `json:"newEmail" validate:"email,nonzero"`
}

// PasswordCreateRequest godoc
type PasswordCreateRequest struct {
	Password string `json:"password" validate:"password"`
}
