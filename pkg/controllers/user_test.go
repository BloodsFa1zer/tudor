package controllers

import (
	"errors"
	"net/http/httptest"
	"strings"
	"study_marketplace/gen/mocks"
	"study_marketplace/pkg/domen/models/entities"
	"study_marketplace/pkg/services"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

var (
	router = gin.New()

	// this is a function that returns a UserController with a mocked UsersRepository
	newTestUserController = func(db *mocks.MockUsersRepository) UserController {
		return NewUsersController(services.NewUserService(
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
	newMockUsersRepository = func(t *testing.T) *mocks.MockUsersRepository {
		ctrl := gomock.NewController(t)
		return mocks.NewMockUsersRepository(ctrl)
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
			`{"email":"john@example.com","name": "John Doe","password": "123456"}`,
			&entities.User{Name: "John Doe", Email: "john@example.com", Verified: true, Role: "user", Password: "123456"},
			&entities.User{1, "John Doe", "john@example.com", "", true, "123456", "user", now, now},
			`{"data":{"token":"token"},"status":"success"}`,
			201,
			nil},
		{"failed_invalid_request",
			`{"email":"`,
			nil,
			nil,
			`{"data":"unexpected EOF","status":"failed"}`,
			400,
			nil},
		{"failed_invalid_email",
			`{"email":"","name": "John Doe","password": "123456"}`,
			nil,
			nil,
			`{"data":"email and password required","status":"failed"}`,
			400,
			nil},
		{"failed_can_not_create_user",
			`{"email": "john@example.com", "name": "John Doe", "password": "123456"}`,
			&entities.User{Name: "John Doe", Email: "john@example.com", Verified: true, Role: "user", Password: "123456"},
			nil,
			`{"data":"can not create user","status":"failed"}`,
			401,
			errors.New("can not create user")},
	}
	for _, tc := range testTable {
		t.Run(tc.scenario, func(t *testing.T) {
			usersRepo := newMockUsersRepository(t)
			usersController := newTestUserController(usersRepo)
			gctx, w := newTestContext("POST", "/api/auth/register", tc.request)

			usersRepo.EXPECT().CreateUser(gomock.Any(), tc.inputDBUser).Return(tc.expectedDBUser, tc.expectedError).AnyTimes()
			usersController.UserRegister(gctx)

			if w.Code != tc.expectedStatusCode {
				t.Errorf("expected status code %d, got %d", tc.expectedStatusCode, w.Code)
			}
			if strings.TrimSpace(w.Body.String()) != tc.expectedResponse {
				t.Errorf("expected response %s, got %s", tc.expectedResponse, w.Body.String())
			}

		})
	}
}
