package server

import (
	config "study_marketplace/pkg/infrastructure/config"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/google"
)

const (
	MaxAge  = 86400 * 30
	IsProde = false
)

func NewServer(conf *config.Config) *gin.Engine {
	s := gin.Default()
	authStore(conf)
	return s
}

func authStore(conf *config.Config) {
	store := sessions.NewCookieStore([]byte(conf.CookieSecret))
	store.MaxAge(MaxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = IsProde

	gothic.Store = store

	goth.UseProviders(
		google.New(conf.GoogleOauthClientId, conf.GoogleOauthClientSecret, conf.GoogleOauthRedirectPage),
		facebook.New(conf.FacebookOauthClientId, conf.FacebookOauthClientSecret, conf.FacebookOauthRedirectPage),
	)
}
