package respmodels

import (
	"time"
)

type SignUpINresponse struct {
	AccessToken string `json:"data"`
	Status      string `json:"status"`
}

type ResponseUser struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Photo     string    `json:"photo"`
	Verified  bool      `json:"verified"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserInfoResponse godoc
type UserInfoResponse struct {
	ResponseUser ResponseUser `json:"data"`
	Status       string       `json:"status"`
}
