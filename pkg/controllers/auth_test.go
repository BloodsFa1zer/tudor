package controllers_test

import (
	"errors"
	"net/http"
	"strings"
	"study_marketplace/gen/mocks"
	"study_marketplace/pkg/controllers"
	"study_marketplace/pkg/domain/models/entities"
	"study_marketplace/pkg/services"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

var (

	// this is a function that returns a UserController with a mocked UsersRepository
	newTestAuthCtrller = func(db *mocks.MockAuthRepository) controllers.AuthController {
		return controllers.NewAuthController("",
			func(res http.ResponseWriter, req *http.Request) (*entities.User, error) {
				return &entities.User{Name: "test", Email: "test@email.com", Verified: true, Role: "user"}, nil
			},
			services.NewAuthService(
				func(userid int64, userName string) (string, error) { return "token", nil },
				db))
	}

	// this is a function that returns a mock controller and a mock repository
	newMockAuthRepository = func(t *testing.T) (*gomock.Controller, *mocks.MockAuthRepository) {
		ctrl := gomock.NewController(t)
		return ctrl, mocks.NewMockAuthRepository(ctrl)
	}
)

func TestAuthWithProviderCallback(t *testing.T) {
	now := time.Now().Truncate(time.Second)
	tt := []struct {
		scenario           string
		inputUser          *entities.User
		expectedUser       *entities.User
		expectedStatusCode int
		expectedResponse   string
		expectedError      error
	}{
		{"success", &entities.User{Name: "test", Email: "test@email.com", Verified: true, Role: "user"},
			&entities.User{Name: "test", Email: "test@email.com", Verified: true, Role: "user", UpdatedAt: now, CreatedAt: now}, http.StatusFound,
			`Found`, nil},
		{"db_error", &entities.User{Name: "test", Email: "test@email.com", Verified: true, Role: "user"}, nil, http.StatusForbidden, "", errors.New("db_error")},
	}
	for _, tc := range tt {
		t.Run(tc.scenario, func(t *testing.T) {
			ctrl, db := newMockAuthRepository(t)
			defer ctrl.Finish()
			ctrller := newTestAuthCtrller(db)
			ctx, w := newTestContext(http.MethodGet, "/auth/callback/google", "")

			db.EXPECT().CreateorUpdateUser(gomock.Any(), tc.inputUser).Return(tc.expectedUser, tc.expectedError)
			ctrller.AuthWithProviderCallback(ctx)

			if w.Code != tc.expectedStatusCode {
				t.Errorf("expected status code %d, got %d", tc.expectedStatusCode, w.Code)
			}
			if !strings.Contains(w.Body.String(), tc.expectedResponse) {
				t.Errorf("expected response %s, got %s", tc.expectedResponse, w.Body.String())
			}
		})
	}

}
