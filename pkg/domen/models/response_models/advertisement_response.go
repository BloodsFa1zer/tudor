package respmodels

// AdvertisementResponse godoc
type AdvertisementResponse struct {
	ID           int64  `json:"id"`
	Title        string `json:"title"`
	ProviderID   int64  `json:"provider_id"`
	ProviderName string `json:"provider_name"`
	Description  string `json:"description"`
	Attachment   string `json:"attachment"`
	Experience   int32  `json:"experience"`
	CategoryName string `json:"category_name"`
	Time         int32  `json:"time"`
	Price        int32  `json:"price"`
	Format       string `json:"format"`
	Language     string `json:"language"`
	MobilePhone  string `json:"mobile_phone"`
	Email        string `json:"email"`
	Telegram     string `json:"telegram"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

// AdvertisementPaginationResponse godoc
type AdvertisementPaginationResponse struct {
	Advertisements []AdvertisementResponse `json:"advertisements"`
	TotalPages     int                     `json:"total_pages"`
	TotalCount     int                     `json:"total_count"`
	Page           int                     `json:"page"`
	PerPage        int                     `json:"per_page"`
	Offset         int                     `json:"offset"`
}
