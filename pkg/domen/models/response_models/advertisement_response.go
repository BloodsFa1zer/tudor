package respmodels

// AdvertisementResponse is a response model for
// getbyid, advertisement-create, advertisement-patch
// swagger:model
type AdvertisementResponse struct {
	ID          int32  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Attachment  string `json:"attachment"`
	Experience  int32  `json:"experience"`
	CategoryID  int32  `json:"category_id"`
	Time        int32  `json:"time"`
	Price       int32  `json:"price"`
	Format      string `json:"format"`
	Language    string `json:"language"`
	MobilePhone string `json:"mobile_phone"`
	Email       string `json:"email"`
	Telegram    string `json:"telegram"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
