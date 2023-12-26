package entities

type Category struct {
	ID             int32           `json:"id"`
	Name           string          `json:"name"`
	ParentCategory *ParentCategory `json:"parent_category"`
}

type ParentCategory struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}
