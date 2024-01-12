package controllers_test

import (
	"net/http"
	"study_marketplace/pkg/controllers"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	tt := []struct {
		scenario           string
		expectedResponse   string
		expectedStatusCode int
	}{
		{"success", `{"data":"Server up and running.","status":"success"}`, 200},
	}
	for _, tc := range tt {
		t.Run(tc.scenario, func(t *testing.T) {
			ctx, w := newTestContext(http.MethodGet, "/api/", "")
			controllers.HealthCheck(ctx)
			if w.Code != tc.expectedStatusCode {
				t.Errorf("expected %d got %d", tc.expectedStatusCode, w.Code)
			}
			if w.Body.String() != tc.expectedResponse {
				t.Errorf("expected %s got %s", tc.expectedResponse, w.Body.String())
			}
		})
	}
}
