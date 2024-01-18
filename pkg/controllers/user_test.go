package controllers_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"study_marketplace/gen/mocks"
	"study_marketplace/pkg/controllers"
	"study_marketplace/pkg/domain/models/entities"
	"study_marketplace/pkg/services"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

var (
	router = gin.New()

	// this is a function that returns a UserController with a mocked UsersRepository
	newTestUserCtrller = func(db *mocks.MockUsersRepository) controllers.UserController {
		return controllers.NewUsersController(services.NewUserService(
			func(userid int64, userName string) (string, error) { return "token", nil },
			func(password string) string { return password },
			func(password, hash string) error {
				if password == hash {
					return nil
				}
				return errors.New("passwords not equal")
			},
			func(token, to string) error { return nil },
			db))
	}

	// this is a function that returns a UserController with a mocked UsersRepository and a request
	newTestContext = func(method, path, request string) (*gin.Context, *httptest.ResponseRecorder) {
		r := httptest.NewRequest(method, path, strings.NewReader(request))
		r.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		ctx := gin.CreateTestContextOnly(w, router)
		ctx.Request = r
		return ctx, w
	}

	// this is a function that returns a UserController with a mocked UsersRepository and a request
	newMockUsersRepository = func(t *testing.T) (*gomock.Controller, *mocks.MockUsersRepository) {
		ctrl := gomock.NewController(t)
		return ctrl, mocks.NewMockUsersRepository(ctrl)
	}

	checkResponse = func(t *testing.T, w *httptest.ResponseRecorder, expectedStatusCode int, expectedResponse string) {
		if w.Code != expectedStatusCode {
			t.Errorf("expected status code %d, got %d", expectedStatusCode, w.Code)
		}
		if strings.TrimSpace(w.Body.String()) != expectedResponse {
			t.Errorf("expected response %s, got %s", expectedResponse, w.Body.String())
		}
	}
)

func TestUserRegister(t *testing.T) {
	now := time.Now()
	testTable := []struct {
		scenario           string
		request            string
		inputDBUser        *entities.User
		expectedDBUser     *entities.User
		expectedResponse   string
		expectedStatusCode int
		expectedError      error
	}{
		{"success",
			`{"email":"john@example.com","name": "John Doe","password": "123456Bb!"}`,
			&entities.User{Name: "John Doe", Email: "john@example.com", Verified: true, Role: "user", Password: "123456Bb!"},
			&entities.User{ID: 1, Name: "John Doe", Email: "john@example.com", Photo: "", Verified: true,
				Password: "123456Bb!", Role: "user", CreatedAt: now, UpdatedAt: now},
			`{"data":{"token":"token"},"status":"success"}`, 201, nil},
		{"failed_invalid_request", `{"email":"`, nil, nil, `{"data":"unexpected EOF","status":"failed"}`, 400, nil},
		{"failed_invalid_email", `{"email":"","name": "John Doe","password": "123456bB!"}`, nil, nil,
			`{"data":"Email: zero value","status":"failed"}`, 400, nil},
		{"failed_can_not_create_user", `{"email": "john@example.com", "name": "John Doe", "password": "123456Bb!"}`,
			&entities.User{Name: "John Doe", Email: "john@example.com", Verified: true, Role: "user", Password: "123456Bb!"}, nil,
			`{"data":"can not create user","status":"failed"}`, 401, errors.New("can not create user")},
	}
	for _, tc := range testTable {
		t.Run(tc.scenario, func(t *testing.T) {
			ctrl, repo := newMockUsersRepository(t)
			defer ctrl.Finish()
			ctrller := newTestUserCtrller(repo)
			gctx, w := newTestContext("POST", "/api/auth/register", tc.request)

			repo.EXPECT().CreateUser(gomock.Any(), tc.inputDBUser).Return(tc.expectedDBUser, tc.expectedError).AnyTimes()
			ctrller.UserRegister(gctx)

			checkResponse(t, w, tc.expectedStatusCode, tc.expectedResponse)
		})
	}
}

func TestUserLogin(t *testing.T) {
	now := time.Now()
	testTable := []struct {
		scenario           string
		request            string
		inputDBEmail       string
		expectedDBUser     *entities.User
		expectedResponse   string
		expectedStatusCode int
		expectedError      error
	}{
		{"success",
			`{"email":"john@example.com","password": "123456bB!"}`, "john@example.com",
			&entities.User{ID: 1, Name: "John Doe", Email: "john@example.com", Photo: "", Verified: true,
				Password: "123456bB!", Role: "user", CreatedAt: now, UpdatedAt: now},
			`{"data":{"token":"token"},"status":"success"}`, 200, nil},
		{"failed_invalid_request", `{"email":"`, "", nil, `{"data":"unexpected EOF","status":"failed"}`, 400, nil},
		{"failed_invalid_email", `{"email":"","name": "John Doe","password": "123456bB!"}`, "", nil,
			`{"data":"Email: zero value","status":"failed"}`, 400, nil},
		{"failed_can_not_fetch_user", `{"email": "john@example.com", "name": "John Doe", "password": "123456bB!"}`, "john@example.com", nil,
			`{"data":"can not fetch user","status":"failed"}`, 401, errors.New("can not fetch user")},
	}
	for _, tc := range testTable {
		t.Run(tc.scenario, func(t *testing.T) {
			ctrl, repo := newMockUsersRepository(t)
			defer ctrl.Finish()
			controller := newTestUserCtrller(repo)
			ctx, w := newTestContext("POST", "/api/auth/login", tc.request)

			repo.EXPECT().GetUserByEmail(gomock.Any(), tc.inputDBEmail).Return(tc.expectedDBUser, tc.expectedError).AnyTimes()
			controller.UserLogin(ctx)

			checkResponse(t, w, tc.expectedStatusCode, tc.expectedResponse)
		})
	}
}

func TestUserInfo(t *testing.T) {
	now := time.Now().Truncate(time.Second)
	testTable := []struct {
		scenario           string
		inputDBId          int64
		expectedDBUser     *entities.User
		expectedResponse   string
		expectedStatusCode int
		expectedError      error
	}{
		{"success",
			1, &entities.User{ID: 1, Name: "John Doe", Email: "john@example.com", Photo: "",
				Verified: true, Password: "123456", Role: "user", CreatedAt: now, UpdatedAt: now},
			`{"data":{"id":1,"name":"John Doe","email":"john@example.com","photo":"","verified":true,"role":"user","created_at":"` +
				now.Format(time.RFC3339) + `","updated_at":"` + now.Format(time.RFC3339) + `"},"status":"success"}`, 200, nil},
		{"failed_invalid_id", 0, nil, `{"data":"user id error","status":"failed"}`, 400, nil},
		{"failed_can_not_fetch_user", 1, nil, `{"data":"can not fetch user","status":"failed"}`, 400, errors.New("can not fetch user")},
	}
	for _, tc := range testTable {
		t.Run(tc.scenario, func(t *testing.T) {
			ctrl, repo := newMockUsersRepository(t)
			defer ctrl.Finish()

			controller := newTestUserCtrller(repo)
			ctx, w := newTestContext("GET", "/api/userinfo", "")
			ctx.Set("user_id", tc.inputDBId)

			repo.EXPECT().GetUserById(gomock.Any(), tc.inputDBId).Return(tc.expectedDBUser, tc.expectedError).AnyTimes()
			controller.UserInfo(ctx)

			checkResponse(t, w, tc.expectedStatusCode, tc.expectedResponse)
		})
	}
}

func TestUserPatch(t *testing.T) {
	now := time.Now().Truncate(time.Second)
	testTable := []struct {
		scenario           string
		request            string
		inputDBId          int64
		inputDBUser        *entities.User
		expectedDBUser     *entities.User
		expectedResponse   string
		expectedStatusCode int
		expectedError      error
	}{
		{"success",
			`{"name": "John Doe", "email": "john@example.com"}`, 1,
			&entities.User{ID: 1, Name: "John Doe", Email: "john@example.com", Verified: true, Role: "user"},
			&entities.User{ID: 1, Name: "John Doe", Email: "john@example.com", Photo: "", Verified: true,
				Password: "123456Bb!", Role: "user", CreatedAt: now, UpdatedAt: now},
			`{"data":{"token":"token"},"status":"success"}`, 200, nil},
		{"failed_invalid_request", `{"email":"`, 1, nil, nil,
			`{"data":"unexpected EOF","status":"failed"}`, 400, nil},
		{"failed_can_not_update_user", `{"name": "John Doe", "email": "john@example.com"}`, 1,
			&entities.User{ID: 1, Name: "John Doe", Email: "john@example.com", Verified: true, Role: "user"}, nil,
			`{"data":"can not fetch user","status":"failed"}`, 400, errors.New("can not fetch user")},
	}
	for _, tc := range testTable {
		t.Run(tc.scenario, func(t *testing.T) {
			ctrl, repo := newMockUsersRepository(t)
			defer ctrl.Finish()
			controller := newTestUserCtrller(repo)
			ctx, w := newTestContext("GET", "/api/protected/user-patch", tc.request)
			ctx.Set("user_id", tc.inputDBId)

			repo.EXPECT().UpdateUser(gomock.Any(), tc.inputDBUser).Return(tc.expectedDBUser, tc.expectedError).AnyTimes()
			controller.UserPatch(ctx)

			checkResponse(t, w, tc.expectedStatusCode, tc.expectedResponse)
		})
	}
}

func TestPasswordReset(t *testing.T) {
	now := time.Now().Truncate(time.Second)
	testTable := []struct {
		scenario           string
		request            string
		inputDBId          int64
		inputDBEmail       string
		expectedDBUser     *entities.User
		expectedResponse   string
		expectedStatusCode int
		expectedError      error
	}{
		{"success",
			`{"email": "john@example.com"}`, 1, "john@example.com",
			&entities.User{ID: 1, Name: "John Doe", Email: "john@example.com", Photo: "", Verified: true,
				Password: "123456", Role: "user", CreatedAt: now, UpdatedAt: now},
			`{"data":"Password Reset Email Has Been Sent","status":"success"}`, 200, nil},
		{"failed_invalid_request", `{"email":"`, 1, "",
			nil, `{"data":"unexpected EOF","status":"failed"}`, 400, nil},
		{"failed_email_not_found", `{"email":""}`, 1, "",
			nil, `{"data":"Email: zero value","status":"failed"}`, 400, nil},
		{"failed_can_not_fetch_user", `{"email": "john@example.com"}`, 1, "john@example.com",
			&entities.User{ID: 1, Name: "John Doe", Email: "john@example.com", Photo: "", Verified: true,
				Password: "123456", Role: "user", CreatedAt: now, UpdatedAt: now},
			`{"data":"Email not found.","status":"failed"}`, 400, errors.New("can not fetch user")},
	}
	for _, tc := range testTable {
		t.Run(tc.scenario, func(t *testing.T) {
			ctrl, repo := newMockUsersRepository(t)
			defer ctrl.Finish()
			controller := newTestUserCtrller(repo)
			ctx, w := newTestContext(http.MethodPost, "/api/protected/user-patch", tc.request)
			ctx.Set("user_id", tc.inputDBId)

			repo.EXPECT().GetUserByEmail(gomock.Any(), tc.inputDBEmail).Return(tc.expectedDBUser, tc.expectedError).AnyTimes()
			controller.PasswordReset(ctx)
			checkResponse(t, w, tc.expectedStatusCode, tc.expectedResponse)
		})
	}
}

func TestPasswordCreate(t *testing.T) {
	now := time.Now().Truncate(time.Second)
	testTable := []struct {
		scenario           string
		request            string
		inputDBId          int64
		inputDBEmail       *entities.User
		expectedDBUser     *entities.User
		expectedResponse   string
		expectedStatusCode int
		expectedError      error
	}{
		{"success",
			`{"password": "123456Bb!"}`, 1, &entities.User{ID: 1, Password: "123456Bb!"},
			&entities.User{ID: 1, Name: "John Doe", Email: "john@example.com", Photo: "", Verified: true,
				Password: "123456Bb!", Role: "user", CreatedAt: now, UpdatedAt: now},
			`{"data":"Password updated.","status":"success"}`, 200, nil},
		{"failed_invalid_request", `{"pass":"`, 1, nil,
			nil, `{"data":"unexpected EOF","status":"failed"}`, 400, nil},
		{"failed_password_not_found", `{"password":""}`, 1, nil,
			nil, `{"data":"Password: password must be at least 8 characters long","status":"failed"}`, 400, nil},
		{"failed_can_not_update_users_password", `{"password": "123456Bb!"}`, 1, &entities.User{ID: 1, Password: "123456Bb!"},
			&entities.User{ID: 1, Name: "John Doe", Email: "john@example.com", Photo: "", Verified: true,
				Password: "123456Bb!", Role: "user", CreatedAt: now, UpdatedAt: now},
			`{"data":"failed request to DB","status":"failed"}`, 400, errors.New("can not fetch user")},
	}
	for _, tc := range testTable {
		t.Run(tc.scenario, func(t *testing.T) {
			ctrl, repo := newMockUsersRepository(t)
			defer ctrl.Finish()
			controller := newTestUserCtrller(repo)

			ctx, w := newTestContext(http.MethodPatch, "/api/auth/reset-password", tc.request)
			ctx.Set("user_id", tc.inputDBId)

			repo.EXPECT().UpdateUser(gomock.Any(), tc.inputDBEmail).Return(tc.expectedDBUser, tc.expectedError).AnyTimes()
			controller.PasswordCreate(ctx)

			checkResponse(t, w, tc.expectedStatusCode, tc.expectedResponse)
		})
	}
}

func TestPasswordChange(t *testing.T) {
	now := time.Now().Truncate(time.Second)
	testTable := []struct {
		scenario           string
		request            string
		inputDBId          int64
		inputDBUser        *entities.User
		expectedDBUser     *entities.User
		expectedResponse   string
		expectedStatusCode int
		expectedError      error
	}{
		{"success",
			`{"currentPassword": "123456bB!", "newPassword": "654321bB!"}`, 1, &entities.User{ID: 1, Password: "654321bB!"},
			&entities.User{ID: 1, Name: "John Doe", Email: "john@example.com", Photo: "", Verified: true,
				Password: "123456bB!", Role: "user", CreatedAt: now, UpdatedAt: now},
			`{"data":"Password has been updated","status":"success"}`, 200, nil},
		{"failed_invalid_request", `{"currentPassword":"`, 1, nil,
			nil, `{"data":"Unable to read the request.","status":"failed"}`, 400, nil},
		{"failed_equal_passwords", `{"currentPassword":"123456bB!", "newPassword": "123456bB!"}`, 1, nil,
			nil, `{"data":"current password and new password are equal","status":"failed"}`, 400, nil},
		{"failed_wrong_password", `{"currentPassword":"111111Bb!", "newPassword": "654321bB!"}`, 1, nil,
			&entities.User{ID: 1, Name: "John Doe", Email: "john@example.com", Photo: "", Verified: true,
				Password: "123456bB!", Role: "user", CreatedAt: now, UpdatedAt: now}, `{"data":"Password change failed: current password is wrong","status":"failed"}`, 401, nil},
	}
	for _, tc := range testTable {
		t.Run(tc.scenario, func(t *testing.T) {
			ctrl, repo := newMockUsersRepository(t)
			defer ctrl.Finish()
			controller := newTestUserCtrller(repo)

			ctx, w := newTestContext(http.MethodPost, "/api/protected/change-password", tc.request)
			ctx.Set("user_id", tc.inputDBId)

			repo.EXPECT().GetUserById(gomock.Any(), tc.inputDBId).Return(tc.expectedDBUser, tc.expectedError).AnyTimes()
			repo.EXPECT().UpdateUser(gomock.Any(), tc.inputDBUser).Return(tc.expectedDBUser, tc.expectedError).AnyTimes()
			controller.PasswordChange(ctx)

			checkResponse(t, w, tc.expectedStatusCode, tc.expectedResponse)
		})
	}
}

func TestEmailChange(t *testing.T) {
	now := time.Now().Truncate(time.Second)
	testTable := []struct {
		scenario           string
		request            string
		inputDBId          int64
		inputDBUser        *entities.User
		expectedDBUser     *entities.User
		expectedResponse   string
		expectedStatusCode int
		expectedError      error
	}{
		{"success",
			`{"currentPassword": "123456bB!", "newEmail": "jane@example.com"}`, 1, &entities.User{ID: 1, Email: "jane@example.com"},
			&entities.User{ID: 1, Name: "John Doe", Email: "john@example.com", Photo: "", Verified: true,
				Password: "123456bB!", Role: "user", CreatedAt: now, UpdatedAt: now},
			`{"data":"Email has been updated","status":"success"}`, 200, nil},
		{"failed_invalid_request", `{"currentPassword":"`, 1, nil,
			nil, `{"data":"Unable to read the request.","status":"failed"}`, 400, nil},
		{"failed_wrong_request", `{"currentPassword":"", "newEmail": ""}`, 1, nil,
			nil, `{"data":"CurrentPassword: password must be at least 8 characters long, NewEmail: zero value","status":"failed"}`, 400, nil},
		{"failed_equal_emails", `{"currentPassword":"123456bB!", "newEmail": "john@example.com"}`, 1, nil,
			&entities.User{ID: 1, Name: "John Doe", Email: "john@example.com", Photo: "", Verified: true,
				Password: "123456bB!", Role: "user", CreatedAt: now, UpdatedAt: now}, `{"data":"Email change failed: current email and new email are equal","status":"failed"}`, 401, nil},
		{"failed_wrong_password", `{"currentPassword":"654321bB!", "newEmail": "jane@example.com"}`, 1, nil,
			&entities.User{ID: 1, Name: "John Doe", Email: "john@example.com", Photo: "", Verified: true,
				Password: "123456bB!", Role: "user", CreatedAt: now, UpdatedAt: now}, `{"data":"Email change failed: current password is wrong","status":"failed"}`, 401, nil},
	}
	for _, tc := range testTable {
		t.Run(tc.scenario, func(t *testing.T) {
			ctrl, repo := newMockUsersRepository(t)
			defer ctrl.Finish()
			controller := newTestUserCtrller(repo)

			ctx, w := newTestContext(http.MethodPost, "/api/protected/change-email", tc.request)
			ctx.Set("user_id", tc.inputDBId)

			repo.EXPECT().GetUserById(gomock.Any(), tc.inputDBId).Return(tc.expectedDBUser, tc.expectedError).AnyTimes()
			repo.EXPECT().UpdateUser(gomock.Any(), tc.inputDBUser).Return(tc.expectedDBUser, tc.expectedError).AnyTimes()
			controller.EmailChange(ctx)

			checkResponse(t, w, tc.expectedStatusCode, tc.expectedResponse)
		})
	}
}
