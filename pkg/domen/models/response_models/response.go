package respmodels

const (
	success = "success"
	failed  = "failed"
)

type response struct {
	Data   interface{} `json:"data"`
	Status string      `json:"status"`
}

func NewResponseFailed(data ...interface{}) (r response) {
	r = response{Data: data, Status: failed}
	return r
}

func NewResponseSuccess(data ...interface{}) (r response) {
	r = response{Data: data, Status: success}
	return r
}
