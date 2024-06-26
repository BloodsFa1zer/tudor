package reqmodels

// CreateAdvertisementRequest godoc
type CreateAdvertisementRequest struct {
	Title        string `json:"title" validate:"min=2,max=50, nonzero"`
	Attachment   string `json:"attachment"`
	Experience   int32  `json:"experience" `
	CategoryName string `json:"category"`
	Time         int32  `json:"time"`
	Price        int32  `json:"price"`
	Currency     string `json:"currency" validate:"advertisementCurrency,nonzero"`
	Format       string `json:"format" validate:"advertisementFormat,nonzero"`
	Language     string `json:"language"`
	Description  string `json:"description"`
	MobilePhone  string `json:"mobile_phone" validate:"phone,nonzero"`
	Email        string `json:"email" validate:"email,nonzero"`
	Telegram     string `json:"telegram"`
}

// UpdateAdvertisementRequest godoc
type UpdateAdvertisementRequest struct {
	ID           int64  `json:"id"`
	Title        string `json:"title" validate:"max=50"`
	Attachment   string `json:"attachment"`
	Experience   int32  `json:"experience"`
	CategoryName string `json:"category"`
	Time         int32  `json:"time"`
	Price        int32  `json:"price"`
	Currency     string `json:"currency" validate:"advertisementCurrency"`
	Format       string `json:"format" validate:"advertisementFormat"`
	Language     string `json:"language"`
	Description  string `json:"description"`
	MobilePhone  string `json:"mobile_phone" validate:"phone"`
	Email        string `json:"email" validate:"email"`
	Telegram     string `json:"telegram"`
}

// DeleteAdvertisementRequest godoc
type DeleteAdvertisementRequest struct {
	ID int64 `json:"id"`
}

// AdvGetFiltered godoc
type AdvertisementFilterRequest struct {
	OrderBy      string `json:"sort_by" validate:"advertisementSortOrder"`
	SortOrder    string `json:"sort_order" validate:"sortOrder"`
	Page         int32  `json:"page"`
	LimitAdv     int32  `json:"per_page"`
	Category     string `json:"category"`
	TimeLength   int32  `json:"time_length"`
	Currency     string `json:"currency" validate:"advertisementCurrency"`
	Format       string `json:"format" validate:"advertisementFormat"`
	MinExp       int32  `json:"min_exp"`
	MaxExp       int32  `json:"max_exp"`
	MinPrice     int32  `json:"min_price"`
	MaxPrice     int32  `json:"max_price"`
	Language     string `json:"language"`
	TitleKeyword string `json:"title_keyword"`
}
