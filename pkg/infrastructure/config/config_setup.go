package configs

import (
	"log"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	DatabaseUrl    string `env:"DB_URL,required"`
	ServerHostname string `env:"B_PORT,required"`
	// Google Oauth Configs
	GoogleOauthClientId     string `env:"GOOGLE_OAUTH_CLIENT_ID,required"`
	GoogleOauthClientSecret string `env:"GOOGLE_OAUTH_CLIENT_SECRET,required"`
	GoogleOauthRedirectPage string `env:"GOOGLE_OAUTH_REDIRECT_PAGE,required"`
	// Facebook Oauth Configs
	FacebookOauthClientId     string `env:"FACEBOOK_OAUTH_CLIENT_ID,required"`
	FacebookOauthClientSecret string `env:"FACEBOOK_OAUTH_CLIENT_SECRET,required"`
	FacebookOauthRedirectPage string `env:"FACEBOOK_OAUTH_REDIRECT_PAGE,required"`
	// Email Configs
	GoogleEmailAddress    string `env:"GOOGLE_EMAIL_ADDRESS"`
	GoogleEmailSecret     string `env:"GOOGLE_EMAIL_SECRET"`
	GoogleEmailSenderName string `env:"GOOGLE_EMAIL_SENDER_NAME"`
	// CORS Configs
	AllowedOrigins []string `env:"ALLOWED_ORIGINS,required" envSeparator:","`
	// JWT Configs
	JWTSecret    string `env:"JWT_SECRET,required"`
	CookieSecret string `env:"COOKIE_SECRET,required"`
	RedirectUrl  string `env:"REDIRECT_URL,required"`
	BasicAppUrl  string `env:"BASIC_APP_URL,required"`
}

func SetUpConfig() *Config {
	var config Config
	if err := env.Parse(&config); err != nil {
		log.Fatalf("env.Parse() in config failed. Error:'%v'", err)
	}
	return &config
}
