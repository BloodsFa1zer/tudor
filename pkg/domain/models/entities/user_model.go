package entities

import (
	"time"
)

// User godoc
type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name" validate:"min=3,max=50"`
	Email     string    `json:"email" validate:"email"`
	Photo     string    `json:"photo"`
	Verified  bool      `json:"verified"`
	Password  string    `json:"password" validate:"password"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
