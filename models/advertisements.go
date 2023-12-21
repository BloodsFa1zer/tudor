package models

type AdvertisementInput struct {
	Title       string `json:"title"`
	Attachment  string `json:"attachment"`
	Experience  int32  `json:"experience"`
	Category    string `json:"category"`
	Time        int32  `json:"time"`
	Price       int32  `json:"price"`
	Format      string `json:"format"`
	Language    string `json:"language"`
	Description string `json:"description"`
	MobilePhone string `json:"mobile_phone"`
	Telegram    string `json:"telegram"`
}

type AdvertisementUpdate struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Attachment  string `json:"attachment"`
	Experience  int32  `json:"experience"`
	Category    string `json:"category"`
	Time        int32  `json:"time"`
	Price       int32  `json:"price"`
	Format      string `json:"format"`
	Language    string `json:"language"`
	Description string `json:"description"`
	MobilePhone string `json:"mobile_phone"`
	Telegram    string `json:"telegram"`
}

type AdvertisementID struct {
	ID int64 `json:"id"`
}

type AdvertisementFilter struct {
	Orderby      string `json:"orderby"`
	Sortorder    string `json:"sortorder"`
	Offsetadv    int32  `json:"offsetadv"`
	Limitadv     int32  `json:"limitadv"`
	Category     string `json:"category"`
	Timelength   int32  `json:"timelength"`
	Format       string `json:"format"`
	Minexp       int32  `json:"minexp"`
	Maxexp       int32  `json:"maxexp"`
	Minprice     int32  `json:"minprice"`
	Maxprice     int32  `json:"maxprice"`
	Language     string `json:"language"`
	Titlekeyword string `json:"titlekeyword"`
}
