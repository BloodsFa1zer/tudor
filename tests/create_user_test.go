package tests

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

var ctx, _ = gin.CreateTestContext(httptest.NewRecorder())

func TestCreateUserDbError(t *testing.T) {
	// errExpected := "db search error"
	// inputModel := queries.User{}
	// want := queries.User{}
	// testRepository := tests_reposotiries.NewTestUsersRepository()
	// testRepository.ReturnIsUserEmailExist = tests_reposotiries.IsUserEmailExistModel{
	// 	Exist: false,
	// 	Error: errors.New("somthing wrong"),
	// }

	// service := services.NewUserService()
	// got, err := service.UserRegister(ctx, inputModel)

	// if err != nil && err.Error() != errExpected {
	// 	t.Errorf("not expected error: [%+v], expected: [%s]", err, errExpected)
	// }

	// if got != want {
	// 	t.Errorf("not expected user: [%+v], expected: [%+v]", got, want)
	// }
}

func TestCreateUserIfEmailAlreadyEsists(t *testing.T) {
	// errExpected := "user with such email already registred"
	// inputModel := queries.User{}
	// want := queries.User{}
	// testRepository := tests_reposotiries.NewTestUsersRepository()
	// testRepository.ReturnIsUserEmailExist = tests_reposotiries.IsUserEmailExistModel{
	// 	Exist: true,
	// 	Error: nil,
	// }
	// SetUsersRepository(testRepository)
	// service := services.NewUserService()
	// got, err := service.UserRegister(ctx, inputModel)

	// if err != nil && err.Error() != errExpected {
	// 	t.Errorf("not expected error: [%+v], expected: [%s]", err, errExpected)
	// }

	// if want == got {
	// 	t.Log("OK check expect user model")
	// }

	// if err == nil {
	// 	t.Error("expected error, but nil returned")
	// }

	// if want != got {
	// 	t.Errorf("not expected user: [%+v], expected: [%+v]", got, want)
	// }
}
