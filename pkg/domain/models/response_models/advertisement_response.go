package respmodels

import (
	"study_marketplace/pkg/domain/models/entities"
	"time"
)

// ResponseAdvertismet godoc
type ResponseAdvertismet struct {
	ID           int64                          `json:"id"`
	Title        string                         `json:"title"`
	ProviderID   int64                          `json:"provider_id"`
	ProviderName string                         `json:"provider_name"`
	Description  string                         `json:"description"`
	Attachment   string                         `json:"attachment"`
	Experience   int32                          `json:"experience"`
	CategoryName string                         `json:"category_name"`
	Time         int32                          `json:"time"`
	Price        int32                          `json:"price"`
	Currency     entities.AdvertisementCurrency `json:"currency"`
	Format       entities.AdvertisementFormat   `json:"format"`
	Language     string                         `json:"language"`
	MobilePhone  string                         `json:"mobile_phone"`
	Email        string                         `json:"email"`
	Telegram     string                         `json:"telegram"`
	CreatedAt    time.Time                      `json:"created_at"`
	UpdatedAt    time.Time                      `json:"updated_at"`
}

// AdvertisementResponse godoc
type AdvertisementResponse struct {
	Advertisement ResponseAdvertismet `json:"data"`
	Status        string              `json:"status"`
}

// AdvertisementsResponse godoc
type AdvertisementsResponse struct {
	Advertisements []ResponseAdvertismet `json:"data"`
	Status         string                `json:"status"`
}

// AdvertisementPaginationResponse godoc
type AdvertisementPaginationResponse struct {
	ResponseAdvertismetPagin ResponseAdvertismetPagin `json:"data"`
	Status                   string                   `json:"status"`
}

// ResponseAdvertismetPagin godoc
type ResponseAdvertismetPagin struct {
	Advertisements []ResponseAdvertismet   `json:"advertisements"`
	PaginationInfo entities.PaginationInfo `json:"pagination_info"`
}
