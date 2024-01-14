package controllers_test

import (
	"study_marketplace/gen/mocks"
	"study_marketplace/pkg/controllers"
	"study_marketplace/pkg/services"
	"testing"

	"github.com/golang/mock/gomock"
)

var (

	// this is a function that returns a UserController with a mocked UsersRepository
	newTestAuthCtrller = func(db *mocks.MockAuthRepository) controllers.AuthController {
		return controllers.NewAuthController("", services.NewAuthService(
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
	// TODO
}
