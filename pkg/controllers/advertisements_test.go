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
			`"price": 100,"currency":"EUR","format":"online","language":"English","description":"test_description","mobile_phone":"0930123456",` +
			`"email":"test@email.com","telegram":"test_telegram"}`, 1,
			&entities.Advertisement{Title: "test", Provider: &entities.User{ID: 1}, Attachment: "test_attachment", Experience: 100500,
				Category: &entities.Category{Name: "English"}, Time: 2, Price: 100, Currency: "EUR", Format: "online", Language: "English",
				Description: "test_description", MobilePhone: "0930123456", Email: "test@email.com", Telegram: "test_telegram"},
			&entities.Advertisement{ID: 1, Title: "test", Provider: &entities.User{ID: 1}, Attachment: "test_attachment", Experience: 10050,
				Category: &entities.Category{Name: "English", ParentCategory: &entities.ParentCategory{Name: "Language learning"}},
				Time:     2, Price: 100, Currency: entities.AdvertisementCurrencyEUR, Format: entities.AdvertisementFormatOnline, Language: "English", Description: "test_description",
				MobilePhone: "0930123456", Email: "test@email.com", Telegram: "test_telegram", CreatedAt: now, UpdatedAt: now},
			`{"data":{"id":1,"title":"test","provider_id":1,"provider_name":"","description":"test_description","attachment":"test_attachment",` +
				`"experience":10050,"category_name":"Language learning: English","time":2,"price":100,"currency":"EUR","format":"online","language":"English",` +
				`"mobile_phone":"0930123456","email":"test@email.com","telegram":"test_telegram","created_at":"` + now.Format(time.RFC3339) + `",` +
				`"updated_at":"` + now.Format(time.RFC3339) + `"},"status":"success"}`,
			http.StatusOK, nil},
		{"phone_validation_failed", `{"title":"test","attachment":"test_attachment","experience":100500,"category":"English","time": 2,` +
			`"price": 100,"currency":"EUR","format":"online","language":"English","description":"test_description","mobile_phone":"wrong",` +
			`"email":"test@email.com","telegram":"test_telegram"}`, 0, nil, nil, `{"data":"MobilePhone: invalid phone number format","status":"failed"}`, http.StatusBadRequest, nil},
		{"currency_validation_failed", `{"title":"test","attachment":"test_attachment","experience":100500,"category":"English","time": 2,` +
			`"price": 100,"currency":"ANY","format":"online","language":"English","description":"test_description","mobile_phone":"0930123456",` +
			`"email":"test@email.com","telegram":"test_telegram"}`, 0, nil, nil, `{"data":"Currency: invalid advertisement currency value","status":"failed"}`, http.StatusBadRequest, nil},
		{"failed_bind_json", `{"title":"test","attachment":"test_attachment","experience":100500,`, 1, nil, nil,
			`{"data":"unexpected EOF","status":"failed"}`, http.StatusBadRequest, nil},
		{"failed_db", `{"title":"test","attachment":"test_attachment","experience":100500,"category":"English","time": 2,"price": 100,"currency":"EUR",` +
			`"format":"online","language":"English","description":"test_description","mobile_phone":"0930123456","email":"test@email.com",` +
			`"telegram":"test_telegram"}`, 1,
			&entities.Advertisement{Title: "test", Provider: &entities.User{ID: 1}, Attachment: "test_attachment", Experience: 100500,
				Category: &entities.Category{Name: "English"}, Time: 2, Price: 100, Currency: "EUR", Format: "online", Language: "English",
				Description: "test_description", MobilePhone: "0930123456", Email: "test@email.com",
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
				Time:     2, Price: 100, Currency: "EUR", Format: "online", Language: "English", Description: "test_description",
				MobilePhone: "0930123456", Email: "test_email", Telegram: "test_telegram", CreatedAt: now, UpdatedAt: now},
			`{"data":{"id":1,"title":"test","provider_id":1,"provider_name":"","description":"test_description","attachment":"test_attachment",` +
				`"experience":10050,"category_name":"Language learning: English","time":2,"price":100,"currency":"EUR","format":"online","language":"English",` +
				`"mobile_phone":"0930123456","email":"test_email","telegram":"test_telegram","created_at":"` + now.Format(time.RFC3339) + `",` +
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

func TestAdvGetAll(t *testing.T) {
	now := time.Now().Truncate(time.Second)
	tt := []struct {
		scenario               string
		expectedAdvertisements []entities.Advertisement
		expectedResponse       string
		expectedStatusCode     int
		expectedError          error
	}{
		{"success",
			genAdvs(now),
			`{"data":[{"id":1,"title":"test","provider_id":1,"provider_name":"John","description":"test_description","attachment":"test_attachment",` +
				`"experience":10050,"category_name":"Language learning: English","time":2,"price":100,"currency":"EUR","format":"online","language":"English",` +
				`"mobile_phone":"test_mobile_phone","email":"test_email","telegram":"test_telegram","created_at":"` + now.Format(time.RFC3339) + `",` +
				`"updated_at":"` + now.Format(time.RFC3339) + `"},{"id":2,"title":"test","provider_id":1,"provider_name":"John","description":"test_description",` +
				`"attachment":"test_attachment","experience":10050,"category_name":"Language learning: English","time":2,"price":100,"currency":"EUR","format":"online",` +
				`"language":"English","mobile_phone":"test_mobile_phone","email":"test_email","telegram":"test_telegram","created_at":"` +
				now.Format(time.RFC3339) + `","updated_at":"` + now.Format(time.RFC3339) + `"}],"status":"success"}`,
			http.StatusOK, nil},
		{"failed_db", nil, `{"data":"db error","status":"failed"}`, http.StatusBadRequest, errors.New("db error")},
	}
	for _, tc := range tt {
		t.Run(tc.scenario, func(t *testing.T) {
			ctrl, db := newMockAdvertisementsRepository(t)
			defer ctrl.Finish()
			ctrller := newTestAdvertisementsCtrller(db)
			ctx, w := newTestContext(http.MethodGet, "/api/open/advertisements/getal", "")

			db.EXPECT().GetAdvertisementAll(gomock.Any()).Return(tc.expectedAdvertisements, tc.expectedError).AnyTimes()
			ctrller.AdvGetAll(ctx)

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
		{"success", "1", 1,
			&entities.Advertisement{ID: 1, Title: "test", Provider: &entities.User{ID: 1}, Attachment: "test_attachment", Experience: 10050,
				Category: &entities.Category{Name: "English", ParentCategory: &entities.ParentCategory{Name: "Language learning"}},
				Time:     2, Price: 100, Currency: "EUR", Format: "online", Language: "English", Description: "test_description",
				MobilePhone: "test_mobile_phone", Email: "test_email", Telegram: "test_telegram", CreatedAt: now, UpdatedAt: now},
			`{"data":{"id":1,"title":"test","provider_id":1,"provider_name":"","description":"test_description","attachment":"test_attachment",` +
				`"experience":10050,"category_name":"Language learning: English","time":2,"price":100,"currency":"EUR","format":"online","language":"English",` +
				`"mobile_phone":"test_mobile_phone","email":"test_email","telegram":"test_telegram","created_at":"` + now.Format(time.RFC3339) +
				`","updated_at":"` + now.Format(time.RFC3339) + `"},"status":"success"}`,
			http.StatusOK, nil},
		{"failed_parce_id", "test", 0, nil, `{"data":"strconv.ParseInt: parsing \"test\": invalid syntax","status":"failed"}`,
			http.StatusBadRequest, nil},
		{"failed_db", "1", 1, nil, `{"data":"db error","status":"failed"}`, http.StatusBadRequest, errors.New("db error")},
	}
	for _, tc := range tt {
		t.Run(tc.scenario, func(t *testing.T) {
			ctrl, db := newMockAdvertisementsRepository(t)
			defer ctrl.Finish()
			ctrller := newTestAdvertisementsCtrller(db)
			ctx, w := newTestContext(http.MethodGet, "/api/open/advertisements/getbyid/:id", "")
			ctx.AddParam("id", tc.inputDBIdStr)

			db.EXPECT().GetAdvertisementByID(gomock.Any(), tc.inputDBIdInt).Return(tc.expectedAdvertisement, tc.expectedError).AnyTimes()
			ctrller.AdvGetByID(ctx)
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
			&reqmodels.AdvertisementFilterRequest{Category: "English", Page: 1, Limitadv: 2},
			&entities.AdvertisementPagination{
				Advertisements: genAdvs(now),
				PaginationInfo: entities.PaginationInfo{TotalPages: 1, TotalCount: 2, Page: 1, PerPage: 2, Offset: 0,
					Orderby: "date", Sortorder: "asc",
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
		{"failed_db", `{"category":"English", "per_page": 2, "page": 1}`, &reqmodels.AdvertisementFilterRequest{Category: "English", Page: 1, Limitadv: 2},
			nil, `{"data":"db error","status":"failed"}`, http.StatusBadRequest, errors.New("db error")},
	}
	for _, tc := range tt {
		t.Run(tc.scenario, func(t *testing.T) {
			ctrl, db := newMockAdvertisementsRepository(t)
			defer ctrl.Finish()
			ctrller := newTestAdvertisementsCtrller(db)
			ctx, w := newTestContext(http.MethodPost, "/api/open/advertisements/adv-filter", tc.request)

			db.EXPECT().FilterAdvertisements(gomock.Any(), tc.inputFilter).Return(tc.expectedAdvertisements, tc.expectedError).AnyTimes()
			ctrller.AdvGetFiltered(ctx)

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
		{"success", 1, genAdvs(now),
			`{"data":[{"id":1,"title":"test","provider_id":1,"provider_name":"John","description":"test_description","attachment":` +
				`"test_attachment","experience":10050,"category_name":"Language learning: English","time":2,"price":100,"currency":"EUR","format":"online",` +
				`"language":"English","mobile_phone":"test_mobile_phone","email":"test_email","telegram":"test_telegram",` +
				`"created_at":"` + now.Format(time.RFC3339) + `","updated_at":"` + now.Format(time.RFC3339) + `"},{"id":2,"title":"test",` +
				`"provider_id":1,"provider_name":"John","description":"test_description","attachment":"test_attachment","experience":10050,` +
				`"category_name":"Language learning: English","time":2,"price":100,"currency":"EUR","format":"online","language":"English",` +
				`"mobile_phone":"test_mobile_phone","email":"test_email","telegram":"test_telegram","created_at":"` + now.Format(time.RFC3339) +
				`","updated_at":"` + now.Format(time.RFC3339) + `"}],"status":"success"}`,
			http.StatusOK, nil},
		{"failed_user_id", 0, nil, `{"data":"Unauthorized.","status":"failed"}`, http.StatusBadRequest, nil},
		{"failed_db", 1, nil, `{"data":"db error","status":"failed"}`, http.StatusBadRequest, errors.New("db error")},
	}
	for _, tc := range tt {
		t.Run(tc.scenario, func(t *testing.T) {
			ctrl, db := newMockAdvertisementsRepository(t)
			defer ctrl.Finish()
			ctrller := newTestAdvertisementsCtrller(db)
			ctx, w := newTestContext(http.MethodGet, "/api/protected/advertisement-getmy", "")
			ctx.Set("user_id", tc.inputUserId)

			db.EXPECT().GetAdvertisementMy(gomock.Any(), tc.inputUserId).Return(tc.expectedAdvertisements, tc.expectedError).AnyTimes()
			ctrller.AdvGetMy(ctx)

			checkResponse(t, w, tc.expectedStatusCode, tc.expectedResponse)
		})
	}
}
