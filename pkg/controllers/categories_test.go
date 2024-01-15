package controllers_test

import (
	"errors"
	"net/http"
	"study_marketplace/database/queries"
	"study_marketplace/gen/mocks"
	"study_marketplace/pkg/controllers"
	"study_marketplace/pkg/services"
	"testing"

	"github.com/golang/mock/gomock"
)

var (
	// this is a function that returns a CategoriesController interface
	newTestCategoryCtrller = func(db *mocks.MockCategoriesRepository) controllers.CategoriesController {
		return controllers.NewCatController(services.NewCategoriesService(db))
	}

	// this is a function that returns a gomock controller and a mock CategoriesRepository interface
	newMockCategoriesRepository = func(t *testing.T) (*gomock.Controller, *mocks.MockCategoriesRepository) {
		ctrl := gomock.NewController(t)
		return ctrl, mocks.NewMockCategoriesRepository(ctrl)
	}
)

func TestCatGetAll(t *testing.T) {
	tt := []struct {
		scenario           string
		especterCategories []queries.GetCategoriesWithChildrenRow
		expectedResponse   string
		expectedStatusCode int
		expectedError      error
	}{
		{"success",
			[]queries.GetCategoriesWithChildrenRow{
				{ParentID: 1, ParentName: "test",
					Children: map[string]interface{}{"children": []interface{}{map[string]interface{}{"id": 9, "name": "test_child"}}}},
				{ParentID: 2, ParentName: "test2",
					Children: map[string]interface{}{"children": []interface{}{map[string]interface{}{"id": 10, "name": "test_child2"}}}}},
			`[{"parent_id":1,"parent_name":"test","children":{"children":[{"id":9,"name":"test_child"}]}},` +
				`{"parent_id":2,"parent_name":"test2","children":{"children":[{"id":10,"name":"test_child2"}]}}]`,
			200,
			nil},
		{"failed",
			[]queries.GetCategoriesWithChildrenRow{},
			`{"data":"failed","status":"failed"}`,
			400,
			errors.New("failed")},
	}

	for _, tc := range tt {
		t.Run(tc.scenario, func(t *testing.T) {
			ctrl, db := newMockCategoriesRepository(t)
			defer ctrl.Finish()

			controller := newTestCategoryCtrller(db)
			ctx, w := newTestContext(http.MethodGet, "/open/allcategories", "")

			db.EXPECT().GetCategoriesWithChildren(ctx).Return(tc.especterCategories, tc.expectedError)
			controller.CatGetAll(ctx)

			if w.Code != tc.expectedStatusCode {
				t.Errorf("expected status code %d, got %d", tc.expectedStatusCode, w.Code)
			}
			if w.Body.String() != tc.expectedResponse {
				t.Errorf("expected response %s, got %s", tc.expectedResponse, w.Body.String())
			}
		})
	}
}
