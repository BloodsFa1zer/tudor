package entities

import (
	"time"
)

type Advertisement struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Provider    *User     `json:"provider"`
	Attachment  string    `json:"attachment"`
	Experience  int32     `json:"experience"`
	Category    *Category `json:"category"`
	Time        int32     `json:"time"`
	Price       int32     `json:"price"`
	Format      string    `json:"format"`
	Language    string    `json:"language"`
	Description string    `json:"description"`
	MobilePhone string    `json:"mobile_phone"`
	Email       string    `json:"email"`
	Telegram    string    `json:"telegram"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
