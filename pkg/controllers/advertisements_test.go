package controllers_test

import (
	"errors"
	"net/http"
	"study_marketplace/gen/mocks"
	"study_marketplace/pkg/controllers"
	"study_marketplace/pkg/domain/models/entities"
	reqmodels "study_marketplace/pkg/domain/models/request_models"
	"study_marketplace/pkg/services"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

var (
	// this is a function that returns a AdvertisementsController interface
	newTestAdvertisementsController = func(db *mocks.MockAdvertisementsRepository) controllers.AdvertisementsControllerInterface {
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
		{scenario: "success",
			request: `{"title":"test","attachment":"test_attachment","experience":100500,"category":"English","time": 2,` +
				`"price": 100,"currency":"EUR","format":"online","language":"English","description":"test_description","mobile_phone":"0930123456",` +
				`"email":"test@email.com","telegram":"test_telegram"}`,
			inputDBId: 1,
			inputAdvertisements: &entities.Advertisement{Title: "test", Provider: &entities.User{ID: 1}, Attachment: "test_attachment", Experience: 100500,
				Category: &entities.Category{Name: "English"}, Time: 2, Price: 100, Currency: "EUR", Format: "online", Language: "English",
				Description: "test_description", MobilePhone: "0930123456", Email: "test@email.com", Telegram: "test_telegram"},
			expectedAdvertisement: &entities.Advertisement{ID: 1, Title: "test", Provider: &entities.User{ID: 1}, Attachment: "test_attachment", Experience: 10050,
				Category: &entities.Category{Name: "English", ParentCategory: &entities.ParentCategory{Name: "Language learning"}},
				Time:     2, Price: 100, Currency: entities.AdvertisementCurrencyEUR, Format: entities.AdvertisementFormatOnline, Language: "English", Description: "test_description",
				MobilePhone: "0930123456", Email: "test@email.com", Telegram: "test_telegram", CreatedAt: now, UpdatedAt: now},
			expectedResponse: `{"data":{"id":1,"title":"test","provider_id":1,"provider_name":"","description":"test_description","attachment":"test_attachment",` +
				`"experience":10050,"category_name":"Language learning: English","time":2,"price":100,"currency":"EUR","format":"online","language":"English",` +
				`"mobile_phone":"0930123456","email":"test@email.com","telegram":"test_telegram","created_at":"` + now.Format(time.RFC3339) + `",` +
				`"updated_at":"` + now.Format(time.RFC3339) + `"},"status":"success"}`,
			expectedStatusCode: http.StatusOK,
			expectedError:      nil,
		},
		{scenario: "phone_validation_failed",
			request: `{"title":"test","attachment":"test_attachment","experience":100500,"category":"English","time": 2,` +
				`"price": 100,"currency":"EUR","format":"online","language":"English","description":"test_description","mobile_phone":"wrong",` +
				`"email":"test@email.com","telegram":"test_telegram"}`,
			inputDBId:             0,
			inputAdvertisements:   nil,
			expectedAdvertisement: nil,
			expectedResponse:      `{"data":"MobilePhone: invalid phone number format","status":"failed"}`,
			expectedStatusCode:    http.StatusBadRequest,
			expectedError:         nil,
		},
		{scenario: "currency_validation_failed",
			request: `{"title":"test","attachment":"test_attachment","experience":100500,"category":"English","time": 2,` +
				`"price": 100,"currency":"ANY","format":"online","language":"English","description":"test_description","mobile_phone":"0930123456",` +
				`"email":"test@email.com","telegram":"test_telegram"}`,
			inputDBId:             0,
			inputAdvertisements:   nil,
			expectedAdvertisement: nil,
			expectedResponse:      `{"data":"Currency: invalid advertisement currency value","status":"failed"}`,
			expectedStatusCode:    http.StatusBadRequest,
			expectedError:         nil,
		},
		{scenario: "failed_bind_json",
			request:               `{"title":"test","attachment":"test_attachment","experience":100500,`,
			inputDBId:             1,
			inputAdvertisements:   nil,
			expectedAdvertisement: nil,
			expectedResponse:      `{"data":"unexpected EOF","status":"failed"}`,
			expectedStatusCode:    http.StatusBadRequest,
			expectedError:         nil,
		},
		{scenario: "failed_db",
			request: `{"title":"test","attachment":"test_attachment","experience":100500,"category":"English","time": 2,"price": 100,"currency":"EUR",` +
				`"format":"online","language":"English","description":"test_description","mobile_phone":"0930123456","email":"test@email.com",` +
				`"telegram":"test_telegram"}`,
			inputDBId: 1,
			inputAdvertisements: &entities.Advertisement{Title: "test", Provider: &entities.User{ID: 1}, Attachment: "test_attachment", Experience: 100500,
				Category: &entities.Category{Name: "English"}, Time: 2, Price: 100, Currency: "EUR", Format: "online", Language: "English",
				Description: "test_description", MobilePhone: "0930123456", Email: "test@email.com",
				Telegram: "test_telegram"},
			expectedAdvertisement: nil,
			expectedResponse:      `{"data":"db error","status":"failed"}`,
			expectedStatusCode:    http.StatusBadRequest,
			expectedError:         errors.New("db error"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.scenario, func(t *testing.T) {
			ctrl, db := newMockAdvertisementsRepository(t)
			defer ctrl.Finish()
			controller := newTestAdvertisementsController(db)
			ctx, w := newTestContext(http.MethodPost, "/api/protected/advertisement-create", tc.request)
			ctx.Set("user_id", tc.inputDBId)

			db.EXPECT().CreateAdvertisement(gomock.Any(), tc.inputAdvertisements).Return(tc.expectedAdvertisement, tc.expectedError).AnyTimes()
			controller.AdvCreate(ctx)

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
		{scenario: "success",
			request:   `{"id":1,"title":"test","attachment":"test_attachment","experience":100500,"category":"English"}`,
			inputDBId: 1,
			inputAdvertisements: &entities.Advertisement{ID: 1, Title: "test", Provider: &entities.User{ID: 1}, Attachment: "test_attachment", Experience: 100500,
				Category: &entities.Category{Name: "English"}},
			expectedAdvertisement: &entities.Advertisement{ID: 1, Title: "test", Provider: &entities.User{ID: 1}, Attachment: "test_attachment", Experience: 10050,
				Category: &entities.Category{Name: "English", ParentCategory: &entities.ParentCategory{Name: "Language learning"}},
				Time:     2, Price: 100, Currency: "EUR", Format: "online", Language: "English", Description: "test_description",
				MobilePhone: "0930123456", Email: "test_email", Telegram: "test_telegram", CreatedAt: now, UpdatedAt: now},
			expectedResponse: `{"data":{"id":1,"title":"test","provider_id":1,"provider_name":"","description":"test_description","attachment":"test_attachment",` +
				`"experience":10050,"category_name":"Language learning: English","time":2,"price":100,"currency":"EUR","format":"online","language":"English",` +
				`"mobile_phone":"0930123456","email":"test_email","telegram":"test_telegram","created_at":"` + now.Format(time.RFC3339) + `",` +
				`"updated_at":"` + now.Format(time.RFC3339) + `"},"status":"success"}`,
			expectedStatusCode: http.StatusOK,
			expectedError:      nil,
		},
		{scenario: "failed_bind_json",
			request:               `{"title":"test","attachment":"test_attachment","experience":100500,`,
			inputDBId:             1,
			inputAdvertisements:   nil,
			expectedAdvertisement: nil,
			expectedResponse:      `{"data":"unexpected EOF","status":"failed"}`,
			expectedStatusCode:    http.StatusBadRequest,
			expectedError:         nil,
		},
		{scenario: "failed_db",
			request:   `{"title":"test","attachment":"test_attachment","experience":100500,"category":"English"}`,
			inputDBId: 1,
			inputAdvertisements: &entities.Advertisement{Title: "test", Provider: &entities.User{ID: 1}, Attachment: "test_attachment", Experience: 100500,
				Category: &entities.Category{Name: "English"}},
			expectedAdvertisement: nil,
			expectedResponse:      `{"data":"db error","status":"failed"}`,
			expectedStatusCode:    http.StatusBadRequest,
			expectedError:         errors.New("db error"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.scenario, func(t *testing.T) {
			ctrl, db := newMockAdvertisementsRepository(t)
			defer ctrl.Finish()
			controller := newTestAdvertisementsController(db)
			ctx, w := newTestContext(http.MethodPatch, "/api/protected/advertisement-patch", tc.request)
			ctx.Set("user_id", tc.inputDBId)

			db.EXPECT().UpdateAdvertisement(gomock.Any(), tc.inputAdvertisements).Return(tc.expectedAdvertisement, tc.expectedError).AnyTimes()
			controller.AdvPatch(ctx)

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
		{scenario: "success",
			request:               `{"id":1}`,
			inputUserId:           1,
			inputAdvertisementsID: 1,
			expectedResponse:      `{"data":"Advertisement deleted","status":"success"}`,
			expectedStatusCode:    http.StatusOK,
			expectedError:         nil,
		},
		{scenario: "failed_user_id",
			request:               `{"id":1}`,
			inputUserId:           0,
			inputAdvertisementsID: 1,
			expectedResponse:      `{"data":"user id error","status":"failed"}`,
			expectedStatusCode:    http.StatusBadRequest,
			expectedError:         nil,
		},
		{scenario: "failed_bind_json",
			request:               `{"id":1`,
			inputUserId:           1,
			inputAdvertisementsID: 0,
			expectedResponse:      `{"data":"unexpected EOF","status":"failed"}`,
			expectedStatusCode:    http.StatusBadRequest,
			expectedError:         nil,
		},
		{scenario: "failed_db",
			request:               `{"id":1}`,
			inputUserId:           1,
			inputAdvertisementsID: 1,
			expectedResponse:      `{"data":"db error","status":"failed"}`,
			expectedStatusCode:    http.StatusBadRequest,
			expectedError:         errors.New("db error"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.scenario, func(t *testing.T) {
			ctrl, db := newMockAdvertisementsRepository(t)
			defer ctrl.Finish()
			controller := newTestAdvertisementsController(db)
			ctx, w := newTestContext(http.MethodDelete, "/api/protected/advertisement-delete", tc.request)
			ctx.Set("user_id", tc.inputUserId)

			db.EXPECT().DeleteAdvertisementByID(gomock.Any(), tc.inputAdvertisementsID, tc.inputUserId).Return(tc.expectedError).AnyTimes()
			controller.AdvDelete(ctx)

			checkResponse(t, w, tc.expectedStatusCode, tc.expectedResponse)
		})
	}
}

func TestAdvGetAll(t *testing.T) {
	now := time.Now().Truncate(time.Second)
	tt := []struct {
		scenario               string
		expectedAdvertisements []entities.Advertisement
		expectedResponse       string
		expectedStatusCode     int
		expectedError          error
	}{
		{scenario: "success",
			expectedAdvertisements: genAdvs(now),
			expectedResponse: `{"data":[{"id":1,"title":"test","provider_id":1,"provider_name":"John","description":"test_description","attachment":"test_attachment",` +
				`"experience":10050,"category_name":"Language learning: English","time":2,"price":100,"currency":"EUR","format":"online","language":"English",` +
				`"mobile_phone":"test_mobile_phone","email":"test_email","telegram":"test_telegram","created_at":"` + now.Format(time.RFC3339) + `",` +
				`"updated_at":"` + now.Format(time.RFC3339) + `"},{"id":2,"title":"test","provider_id":1,"provider_name":"John","description":"test_description",` +
				`"attachment":"test_attachment","experience":10050,"category_name":"Language learning: English","time":2,"price":100,"currency":"EUR","format":"online",` +
				`"language":"English","mobile_phone":"test_mobile_phone","email":"test_email","telegram":"test_telegram","created_at":"` +
				now.Format(time.RFC3339) + `","updated_at":"` + now.Format(time.RFC3339) + `"}],"status":"success"}`,
			expectedStatusCode: http.StatusOK,
			expectedError:      nil,
		},
		{scenario: "failed_db",
			expectedAdvertisements: nil,
			expectedResponse:       `{"data":"db error","status":"failed"}`,
			expectedStatusCode:     http.StatusBadRequest,
			expectedError:          errors.New("db error"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.scenario, func(t *testing.T) {
			ctrl, db := newMockAdvertisementsRepository(t)
			defer ctrl.Finish()
			controller := newTestAdvertisementsController(db)
			ctx, w := newTestContext(http.MethodGet, "/api/open/advertisements/getal", "")

			db.EXPECT().GetAdvertisementAll(gomock.Any()).Return(tc.expectedAdvertisements, tc.expectedError).AnyTimes()
			controller.AdvGetAll(ctx)

			checkResponse(t, w, tc.expectedStatusCode, tc.expectedResponse)
		})
	}
}

func genAdvs(now time.Time) []entities.Advertisement {
	advs := make([]entities.Advertisement, 2)
	for i := range advs {
		advs[i] = entities.Advertisement{ID: int64(i + 1), Title: "test", Provider: &entities.User{ID: 1, Name: "John"}, Attachment: "test_attachment", Experience: 10050,
			Category: &entities.Category{Name: "English", ParentCategory: &entities.ParentCategory{Name: "Language learning"}},
			Time:     2, Price: 100, Currency: "EUR", Format: "online", Language: "English", Description: "test_description",
			MobilePhone: "test_mobile_phone", Email: "test_email", Telegram: "test_telegram", CreatedAt: now, UpdatedAt: now}
	}
	return advs
}

func TestAdvGetByID(t *testing.T) {
	now := time.Now().Truncate(time.Second)
	tt := []struct {
		scenario              string
		inputDBIdStr          string
		inputDBIdInt          int64
		expectedAdvertisement *entities.Advertisement
		expectedResponse      string
		expectedStatusCode    int
		expectedError         error
	}{
		{scenario: "success",
			inputDBIdStr: "1",
			inputDBIdInt: 1,
			expectedAdvertisement: &entities.Advertisement{ID: 1, Title: "test", Provider: &entities.User{ID: 1}, Attachment: "test_attachment", Experience: 10050,
				Category: &entities.Category{Name: "English", ParentCategory: &entities.ParentCategory{Name: "Language learning"}},
				Time:     2, Price: 100, Currency: "EUR", Format: "online", Language: "English", Description: "test_description",
				MobilePhone: "test_mobile_phone", Email: "test_email", Telegram: "test_telegram", CreatedAt: now, UpdatedAt: now},
			expectedResponse: `{"data":{"id":1,"title":"test","provider_id":1,"provider_name":"","description":"test_description","attachment":"test_attachment",` +
				`"experience":10050,"category_name":"Language learning: English","time":2,"price":100,"currency":"EUR","format":"online","language":"English",` +
				`"mobile_phone":"test_mobile_phone","email":"test_email","telegram":"test_telegram","created_at":"` + now.Format(time.RFC3339) +
				`","updated_at":"` + now.Format(time.RFC3339) + `"},"status":"success"}`,
			expectedStatusCode: http.StatusOK,
			expectedError:      nil,
		},
		{scenario: "failed_parce_id",
			inputDBIdStr:          "test",
			inputDBIdInt:          0,
			expectedAdvertisement: nil,
			expectedResponse:      `{"data":"strconv.ParseInt: parsing \"test\": invalid syntax","status":"failed"}`,
			expectedStatusCode:    http.StatusBadRequest,
			expectedError:         nil,
		},
		{scenario: "failed_db",
			inputDBIdStr:          "1",
			inputDBIdInt:          1,
			expectedAdvertisement: nil,
			expectedResponse:      `{"data":"db error","status":"failed"}`,
			expectedStatusCode:    http.StatusBadRequest,
			expectedError:         errors.New("db error"),
		},
	}
	for _, tc := range tt {
		t.Run(tc.scenario, func(t *testing.T) {
			ctrl, db := newMockAdvertisementsRepository(t)
			defer ctrl.Finish()
			controller := newTestAdvertisementsController(db)
			ctx, w := newTestContext(http.MethodGet, "/api/open/advertisements/getbyid/:id", "")
			ctx.AddParam("id", tc.inputDBIdStr)

			db.EXPECT().GetAdvertisementByID(gomock.Any(), tc.inputDBIdInt).Return(tc.expectedAdvertisement, tc.expectedError).AnyTimes()
			controller.AdvGetByID(ctx)
			checkResponse(t, w, tc.expectedStatusCode, tc.expectedResponse)
		})
	}
}

func TestAdvGetFiltered(t *testing.T) {
	now := time.Now().Truncate(time.Second)
	tt := []struct {
		scenario               string
		request                string
		inputFilter            *reqmodels.AdvertisementFilterRequest
		expectedAdvertisements *entities.AdvertisementPagination
		expectedResponse       string
		expectedStatusCode     int
		expectedError          error
	}{
		{"success",
			`{"category":"English", "per_page": 2, "page": 1}`,
			&reqmodels.AdvertisementFilterRequest{Category: "English", Page: 1, LimitAdv: 2},
			&entities.AdvertisementPagination{
				Advertisements: genAdvs(now),
				PaginationInfo: entities.PaginationInfo{TotalPages: 1, TotalCount: 2, Page: 1, PerPage: 2, Offset: 0,
					OrderBy: "date", SortOrder: "asc",
				}},
			`{"data":{"advertisements":[{"id":1,"title":"test","provider_id":1,"provider_name":"John","description":"test_description",` +
				`"attachment":"test_attachment","experience":10050,"category_name":"Language learning: English","time":2,"price":100,"currency":"EUR",` +
				`"format":"online","language":"English","mobile_phone":"test_mobile_phone","email":"test_email","telegram":"test_telegram",` +
				`"created_at":"` + now.Format(time.RFC3339) + `","updated_at":"` + now.Format(time.RFC3339) + `"},{"id":2,"title":"test",` +
				`"provider_id":1,"provider_name":"John","description":"test_description","attachment":"test_attachment","experience":10050,` +
				`"category_name":"Language learning: English","time":2,"price":100,"currency":"EUR","format":"online","language":"English",` +
				`"mobile_phone":"test_mobile_phone","email":"test_email","telegram":"test_telegram","created_at":"` +
				now.Format(time.RFC3339) + `","updated_at":"` + now.Format(time.RFC3339) + `"}],"pagination_info":` +
				`{"total_pages":1,"total_count":2,"page":1,"per_page":2,"offset":0,"sort_by":"date","sort_order":"asc"}},"status":"success"}`,
			http.StatusOK, nil},
		{"failed_bind_json", `{"category":"English", "per_page": 2, "page": 1`, nil, nil,
			`{"data":"unexpected EOF","status":"failed"}`, http.StatusBadRequest, nil},
		{"failed_db", `{"category":"English", "per_page": 2, "page": 1}`, &reqmodels.AdvertisementFilterRequest{Category: "English", Page: 1, LimitAdv: 2},
			nil, `{"data":"db error","status":"failed"}`, http.StatusBadRequest, errors.New("db error")},
	}
	for _, tc := range tt {
		t.Run(tc.scenario, func(t *testing.T) {
			ctrl, db := newMockAdvertisementsRepository(t)
			defer ctrl.Finish()
			controller := newTestAdvertisementsController(db)
			ctx, w := newTestContext(http.MethodPost, "/api/open/advertisements/adv-filter", tc.request)

			db.EXPECT().FilterAdvertisements(gomock.Any(), tc.inputFilter).Return(tc.expectedAdvertisements, tc.expectedError).AnyTimes()
			controller.AdvGetFiltered(ctx)

			checkResponse(t, w, tc.expectedStatusCode, tc.expectedResponse)
		})
	}
}

func TestAdvGetMy(t *testing.T) {
	now := time.Now().Truncate(time.Second)
	tt := []struct {
		scenario               string
		inputUserId            int64
		expectedAdvertisements []entities.Advertisement
		expectedResponse       string
		expectedStatusCode     int
		expectedError          error
	}{
		{scenario: "success",
			inputUserId:            1,
			expectedAdvertisements: genAdvs(now),
			expectedResponse: `{"data":[{"id":1,"title":"test","provider_id":1,"provider_name":"John","description":"test_description","attachment":` +
				`"test_attachment","experience":10050,"category_name":"Language learning: English","time":2,"price":100,"currency":"EUR","format":"online",` +
				`"language":"English","mobile_phone":"test_mobile_phone","email":"test_email","telegram":"test_telegram",` +
				`"created_at":"` + now.Format(time.RFC3339) + `","updated_at":"` + now.Format(time.RFC3339) + `"},{"id":2,"title":"test",` +
				`"provider_id":1,"provider_name":"John","description":"test_description","attachment":"test_attachment","experience":10050,` +
				`"category_name":"Language learning: English","time":2,"price":100,"currency":"EUR","format":"online","language":"English",` +
				`"mobile_phone":"test_mobile_phone","email":"test_email","telegram":"test_telegram","created_at":"` + now.Format(time.RFC3339) +
				`","updated_at":"` + now.Format(time.RFC3339) + `"}],"status":"success"}`,
			expectedStatusCode: http.StatusOK,
			expectedError:      nil,
		},
		{scenario: "failed_user_id",
			inputUserId:            0,
			expectedAdvertisements: nil,
			expectedResponse:       `{"data":"Unauthorized.","status":"failed"}`,
			expectedStatusCode:     http.StatusBadRequest,
			expectedError:          nil,
		},
		{scenario: "failed_db",
			inputUserId:            1,
			expectedAdvertisements: nil,
			expectedResponse:       `{"data":"db error","status":"failed"}`,
			expectedStatusCode:     http.StatusBadRequest,
			expectedError:          errors.New("db error"),
		},
	}
	for _, tc := range tt {
		t.Run(tc.scenario, func(t *testing.T) {
			ctrl, db := newMockAdvertisementsRepository(t)
			defer ctrl.Finish()
			controller := newTestAdvertisementsController(db)
			ctx, w := newTestContext(http.MethodGet, "/api/protected/advertisement-getmy", "")
			ctx.Set("user_id", tc.inputUserId)

			db.EXPECT().GetAdvertisementMy(gomock.Any(), tc.inputUserId).Return(tc.expectedAdvertisements, tc.expectedError).AnyTimes()
			controller.AdvGetMy(ctx)

			checkResponse(t, w, tc.expectedStatusCode, tc.expectedResponse)
		})
	}
}
