package models

import "time"

const (
	success = "success"
	failed  = "failed"
)

type response struct {
	Data   interface{} `json:"data"`
	Status string      `json:"status"`
}

func NewResponseFailed(data interface{}) (r response) {
	r = response{Data: data, Status: failed}
	return r
}

func NewResponseSuccess(data interface{}) (r response) {
	r = response{Data: data, Status: success}
	return r
}

type UserResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Photo     string    `json:"photo"`
	Verified  bool      `json:"verified"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
