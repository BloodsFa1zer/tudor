// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package queries

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Advertisement struct {
	ID          int64       `json:"id"`
	Title       string      `json:"title"`
	ProviderID  int64       `json:"provider_id"`
	Attachment  string      `json:"attachment"`
	Experience  int32       `json:"experience"`
	CategoryID  int64       `json:"category_id"`
	Time        int32       `json:"time"`
	Price       int32       `json:"price"`
	Format      string      `json:"format"`
	Language    string      `json:"language"`
	Description string      `json:"description"`
	MobilePhone pgtype.Text `json:"mobile_phone"`
	Email       pgtype.Text `json:"email"`
	Telegram    pgtype.Text `json:"telegram"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

type Category struct {
	ID       int32       `json:"id"`
	Name     string      `json:"name"`
	ParentID pgtype.Int4 `json:"parent_id"`
}

type User struct {
	ID        int64       `json:"id"`
	Name      pgtype.Text `json:"name"`
	Email     string      `json:"email"`
	Photo     pgtype.Text `json:"photo"`
	Verified  bool        `json:"verified"`
	Password  pgtype.Text `json:"password"`
	Role      string      `json:"role"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}
