package controllers_test

import (
	"errors"
	"net/http"
	"study_marketplace/gen/mocks"
	"study_marketplace/pkg/controllers"
	"study_marketplace/pkg/domain/models/entities"
	"study_marketplace/pkg/services"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

var (
	// this is a function that returns a AdvertisementsController interface
	newTestAdvertisementsCtrller = func(db *mocks.MockAdvertisementsRepository) controllers.AdvertisementsController {
		return controllers.NewAdvertisementsController(services.NewAdvertisementService(db))
	}

	// this is a function that returns a gomock controller and a mock AdvertisementsRepository interface
	newMockAdvertisementsRepository = func(t *testing.T) (*gomock.Controller, *mocks.MockAdvertisementsRepository) {
		ctrl := gomock.NewController(t)
		return ctrl, mocks.NewMockAdvertisementsRepository(ctrl)
	}
)

func TestAdvCreate(t *testing.T) {
	now := time.Now().Truncate(time.Second)
	tt := []struct {
		scenario              string
		request               string
		inputDBId             int64
		inputAdvertisements   *entities.Advertisement
		expectedAdvertisement *entities.Advertisement
		expectedResponse      string
		expectedStatusCode    int
		expectedError         error
	}{
		{"success", `{"title":"test","attachment":"test_attachment","experience":100500,"category":"English","time": 2,` +
			`"price": 100,"format":"online","language":"English","description":"test_description","mobile_phone":"test_mobile_phone",` +
			`"email":"test_email","telegram":"test_telegram"}`, 1,
			&entities.Advertisement{Title: "test", Provider: &entities.User{ID: 1}, Attachment: "test_attachment", Experience: 100500,
				Category: &entities.Category{Name: "English"}, Time: 2, Price: 100, Format: "online", Language: "English",
				Description: "test_description", MobilePhone: "test_mobile_phone", Email: "test_email", Telegram: "test_telegram"},
			&entities.Advertisement{ID: 1, Title: "test", Provider: &entities.User{ID: 1}, Attachment: "test_attachment", Experience: 10050,
				Category: &entities.Category{Name: "English", ParentCategory: &entities.ParentCategory{Name: "Language learning"}},
				Time:     2, Price: 100, Format: "online", Language: "English", Description: "test_description",
				MobilePhone: "test_mobile_phone", Email: "test_email", Telegram: "test_telegram", CreatedAt: now, UpdatedAt: now},
			`{"data":{"id":1,"title":"test","provider_id":1,"provider_name":"","description":"test_description","attachment":"test_attachment",` +
				`"experience":10050,"category_name":"Language learning: English","time":2,"price":100,"format":"online","language":"English",` +
				`"mobile_phone":"test_mobile_phone","email":"test_email","telegram":"test_telegram","created_at":"` + now.Format(time.RFC3339) + `",` +
				`"updated_at":"` + now.Format(time.RFC3339) + `"},"status":"success"}`,
			http.StatusOK, nil},
		{"failed_user_id", `{"title":"test","attachment":"test_attachment","experience":100500,"category":"English","time": 2,` +
			`"price": 100,"format":"online","language":"English","description":"test_description","mobile_phone":"test_mobile_phone",` +
			`"email":"test_email","telegram":"test_telegram"}`, 0, nil, nil, `{"data":"user id error","status":"failed"}`, http.StatusBadRequest, nil},
		{"failed_bind_json", `{"title":"test","attachment":"test_attachment","experience":100500,`, 1, nil, nil,
			`{"data":"unexpected EOF","status":"failed"}`, http.StatusBadRequest, nil},
		{"failed_db", `{"title":"test","attachment":"test_attachment","experience":100500,"category":"English","time": 2,"price": 100,` +
			`"format":"online","language":"English","description":"test_description","mobile_phone":"test_mobile_phone","email":"test_email",` +
			`"telegram":"test_telegram"}`, 1,
			&entities.Advertisement{Title: "test", Provider: &entities.User{ID: 1}, Attachment: "test_attachment", Experience: 100500,
				Category: &entities.Category{Name: "English"}, Time: 2, Price: 100, Format: "online", Language: "English",
				Description: "test_description", MobilePhone: "test_mobile_phone", Email: "test_email",
				Telegram: "test_telegram"}, nil, `{"data":"db error","status":"failed"}`, http.StatusBadRequest, errors.New("db error")},
	}
	for _, tc := range tt {
		t.Run(tc.scenario, func(t *testing.T) {
			ctrl, db := newMockAdvertisementsRepository(t)
			defer ctrl.Finish()
			ctrller := newTestAdvertisementsCtrller(db)
			ctx, w := newTestContext(http.MethodPost, "/api/protected/advertisement-create", tc.request)
			ctx.Set("user_id", tc.inputDBId)

			db.EXPECT().CreateAdvertisement(gomock.Any(), tc.inputAdvertisements).Return(tc.expectedAdvertisement, tc.expectedError).AnyTimes()
			ctrller.AdvCreate(ctx)

			checkResponse(t, w, tc.expectedStatusCode, tc.expectedResponse)
		})
	}
}

func TestAdvPatch(t *testing.T) {
	now := time.Now().Truncate(time.Second)
	tt := []struct {
		scenario              string
		request               string
		inputDBId             int64
		inputAdvertisements   *entities.Advertisement
		expectedAdvertisement *entities.Advertisement
		expectedResponse      string
		expectedStatusCode    int
		expectedError         error
	}{
		{"success", `{"id":1,"title":"test","attachment":"test_attachment","experience":100500,"category":"English"}`, 1,
			&entities.Advertisement{ID: 1, Title: "test", Provider: &entities.User{ID: 1}, Attachment: "test_attachment", Experience: 100500,
				Category: &entities.Category{Name: "English"}},
			&entities.Advertisement{ID: 1, Title: "test", Provider: &entities.User{ID: 1}, Attachment: "test_attachment", Experience: 10050,
				Category: &entities.Category{Name: "English", ParentCategory: &entities.ParentCategory{Name: "Language learning"}},
				Time:     2, Price: 100, Format: "online", Language: "English", Description: "test_description",
				MobilePhone: "test_mobile_phone", Email: "test_email", Telegram: "test_telegram", CreatedAt: now, UpdatedAt: now},
			`{"data":{"id":1,"title":"test","provider_id":1,"provider_name":"","description":"test_description","attachment":"test_attachment",` +
				`"experience":10050,"category_name":"Language learning: English","time":2,"price":100,"format":"online","language":"English",` +
				`"mobile_phone":"test_mobile_phone","email":"test_email","telegram":"test_telegram","created_at":"` + now.Format(time.RFC3339) + `",` +
				`"updated_at":"` + now.Format(time.RFC3339) + `"},"status":"success"}`,
			http.StatusOK, nil},
		{"failed_bind_json", `{"title":"test","attachment":"test_attachment","experience":100500,`, 1, nil, nil,
			`{"data":"unexpected EOF","status":"failed"}`, http.StatusBadRequest, nil},
		{"failed_db", `{"title":"test","attachment":"test_attachment","experience":100500,"category":"English"}`, 1,
			&entities.Advertisement{Title: "test", Provider: &entities.User{ID: 1}, Attachment: "test_attachment", Experience: 100500,
				Category: &entities.Category{Name: "English"}},
			nil, `{"data":"db error","status":"failed"}`, http.StatusBadRequest, errors.New("db error")},
	}
	for _, tc := range tt {
		t.Run(tc.scenario, func(t *testing.T) {
			ctrl, db := newMockAdvertisementsRepository(t)
			defer ctrl.Finish()
			ctrller := newTestAdvertisementsCtrller(db)
			ctx, w := newTestContext(http.MethodPatch, "/api/protected/advertisement-patch", tc.request)
			ctx.Set("user_id", tc.inputDBId)

			db.EXPECT().UpdateAdvertisement(gomock.Any(), tc.inputAdvertisements).Return(tc.expectedAdvertisement, tc.expectedError).AnyTimes()
			ctrller.AdvPatch(ctx)

			checkResponse(t, w, tc.expectedStatusCode, tc.expectedResponse)
		})
	}
}

func TestAdvDelete(t *testing.T) {
	// now := time.Now().Truncate(time.Second)
	tt := []struct {
		scenario              string
		request               string
		inputUserId           int64
		inputAdvertisementsID int64
		expectedResponse      string
		expectedStatusCode    int
		expectedError         error
	}{
		{"success", `{"id":1}`, 1, 1, `{"data":"Advertisement deleted","status":"success"}`, http.StatusOK, nil},
		{"failed_user_id", `{"id":1}`, 0, 1, `{"data":"user id error","status":"failed"}`, http.StatusBadRequest, nil},
		{"failed_bind_json", `{"id":1`, 1, 0, `{"data":"unexpected EOF","status":"failed"}`, http.StatusBadRequest, nil},
		{"failed_db", `{"id":1}`, 1, 1, `{"data":"db error","status":"failed"}`, http.StatusBadRequest, errors.New("db error")},
	}
	for _, tc := range tt {
		t.Run(tc.scenario, func(t *testing.T) {
			ctrl, db := newMockAdvertisementsRepository(t)
			defer ctrl.Finish()
			ctrller := newTestAdvertisementsCtrller(db)
			ctx, w := newTestContext(http.MethodDelete, "/api/protected/advertisement-delete", tc.request)
			ctx.Set("user_id", tc.inputUserId)

			db.EXPECT().DeleteAdvertisementByID(gomock.Any(), tc.inputAdvertisementsID, tc.inputUserId).Return(tc.expectedError).AnyTimes()
			ctrller.AdvDelete(ctx)

			checkResponse(t, w, tc.expectedStatusCode, tc.expectedResponse)
		})
	}
}
