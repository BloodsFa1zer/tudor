package reqmodels

type CreateUpdateAdvertisementRequest struct {
	Title        string `json:"title"`
	Attachment   string `json:"attachment"`
	Experience   int32  `json:"experience"`
	CategoryName string `json:"category_name"`
	Price        int32  `json:"price"`
	Format       string `json:"format"`
	Language     string `json:"language"`
	Description  string `json:"description"`
	MobilePhone  string `json:"mobile_phone"`
	Email        string `json:"email"`
	Telegram     string `json:"telegram"`
}
