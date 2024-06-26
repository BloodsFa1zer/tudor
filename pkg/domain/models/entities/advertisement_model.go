package entities

import (
	"time"
)

type AdvertisementFormat string

const (
	AdvertisementFormatOnline  AdvertisementFormat = "online"
	AdvertisementFormatOffline AdvertisementFormat = "offline"
)

type AdvertisementCurrency string

const (
	AdvertisementCurrencyEUR AdvertisementCurrency = "EUR"
	AdvertisementCurrencyUSD AdvertisementCurrency = "USD"
	AdvertisementCurrencyUAH AdvertisementCurrency = "UAH"
)

type Advertisement struct {
	ID          int64                 `json:"id"`
	Title       string                `json:"title"`
	Provider    *User                 `json:"provider"`
	Attachment  string                `json:"attachment"`
	Experience  int                   `json:"experience"`
	Category    *Category             `json:"category"`
	Time        int                   `json:"time"`
	Price       int                   `json:"price"`
	Currency    AdvertisementCurrency `json:"currency"`
	Format      AdvertisementFormat   `json:"format"`
	Language    string                `json:"language"`
	Description string                `json:"description"`
	MobilePhone string                `json:"mobile_phone"`
	Email       string                `json:"email"`
	Telegram    string                `json:"telegram"`
	CreatedAt   time.Time             `json:"created_at"`
	UpdatedAt   time.Time             `json:"updated_at"`
}

type AdvertisementPagination struct {
	Advertisements []Advertisement `json:"advertisements"`
	PaginationInfo PaginationInfo  `json:"pagination_info"`
}

type PaginationInfo struct {
	TotalPages int    `json:"total_pages"`
	TotalCount int    `json:"total_count"`
	Page       int    `json:"page"`
	PerPage    int    `json:"per_page"`
	Offset     int    `json:"offset"`
	OrderBy    string `json:"sort_by"`
	SortOrder  string `json:"sort_order"`
}
