package reqm

import (
	respmodels "study_marketplace/pkg/domain/models/response_models"
)

func FailedResponse(reason string) *respmodels.FailedResponse {
	return &respmodels.FailedResponse{
		Data:   reason,
		Status: "failed",
	}
}

func StrResponse(reason string) *respmodels.StringResponse {
	return &respmodels.StringResponse{
		Data:   reason,
		Status: "success",
	}
}
