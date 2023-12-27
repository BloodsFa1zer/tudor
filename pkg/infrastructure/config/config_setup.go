package configs

import (
	"log"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	DatabaseUrlPool           string   `env:"DATABASE_URL,required"`
	DatabaseUrl               string   `env:"MIGRATION_URL,required"`
	ServerHostname            string   `env:"SERVER_HOSTNAME,required"`
	DocsHostname              string   `env:"DOCS_HOSTNAME,required"`
	GoogleCallbackDomain      string   `env:"GOOGLE_CALLBACK_DOMAIN,required"`
	GoogleOauthClientId       string   `env:"GOOGLE_OAUTH_CLIENT_ID,required"`
	GoogleOauthClientSecret   string   `env:"GOOGLE_OAUTH_CLIENT_SECRET,required"`
	GoogleOauthRedirectPage   string   `env:"GOOGLE_OAUTH_REDIRECT_PAGE,required"`
	FacebookOauthClientId     string   `env:"FACEBOOK_OAUTH_CLIENT_ID,required"`
	FacebookOauthClientSecret string   `env:"FACEBOOK_OAUTH_CLIENT_SECRET,required"`
	FacebookOauthRedirectPage string   `env:"FACEBOOK_OAUTH_REDIRECT_PAGE,required"`
	GoogleEmailAddress        string   `env:"GOOGLE_EMAIL_ADDRESS"`
	GoogleEmailSecret         string   `env:"GOOGLE_EMAIL_SECRET"`
	PasswordResetRedirectPage string   `env:"PASSWORD_RESET_REDIRECT_PAGE"`
	AllowedOrigins            []string `env:"ALLOWED_ORIGINS,required" envSeparator:","`
	JWTSecret                 string   `env:"JWT_SECRET,required"`
	CookieSecret              string   `env:"COOKIE_SECRET,required"`
	RedirectUrl               string   `env:"REDIRECT_URL,required"`
}

func SetUpConfig() *Config {
	var config Config
	if err := env.Parse(&config); err != nil {
		log.Fatalf("env.Parse() in config failed. Error:'%v'", err)
	}
	return &config
}
