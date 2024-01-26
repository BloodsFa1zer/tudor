package entities

type Avatar struct {
	ID       int64  `json:"id"`
	Filename string `json:"filename"`
	Format   string `json:"format"`
	Data     []byte `json:"data"`
	FileAddr string `json:"file_addr"`
	Provider *User  `json:"provider"`
}
