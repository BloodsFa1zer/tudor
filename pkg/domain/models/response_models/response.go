package respmodels

type StringResponse struct {
	Data   string `json:"data"`
	Status string `json:"status"`
}

type FailedResponse struct {
	Data   string `json:"data"`
	Status string `json:"status"`
}
