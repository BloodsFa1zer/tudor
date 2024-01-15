package registry

import (
	"net/http"
	"study_marketplace/pkg/domain/models/entities"

	"github.com/markbates/goth/gothic"
)

func callbackfunc() func(res http.ResponseWriter, req *http.Request) (*entities.User, error) {
	return func(res http.ResponseWriter, req *http.Request) (*entities.User, error) {
		user, err := gothic.CompleteUserAuth(res, req)
		if err != nil {
			return nil, err
		}

		return &entities.User{
			Name:     user.Name,
			Email:    user.Email,
			Photo:    user.AvatarURL,
			Verified: true,
			Role:     "user",
		}, nil
	}
}
