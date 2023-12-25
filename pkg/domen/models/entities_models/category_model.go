package entities

type Category struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	ParentID int32  `json:"parent_id"`
}
