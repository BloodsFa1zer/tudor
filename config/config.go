package configs

import (
	"log"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	DatabseUrl                string   `env:"DATABASE_URL"`
	ServerHostname            string   `env:"SERVER_HOSTNAME"`
	DocsHostname              string   `env:"DOCS_HOSTNAME"`
	GoogleCallbackDomain      string   `env:"GOOGLE_CALLBACK_DOMAIN"`
	GoogleOauthClientId       string   `env:"GOOGLE_OAUTH_CLIENT_ID"`
	GoogleOauthClientSecret   string   `env:"GOOGLE_OAUTH_CLIENT_SECRET"`
	GoogleOauthRedirectPage   string   `env:"GOOGLE_OAUTH_REDIRECT_PAGE"`
	GoogleEmailAddress        string   `env:"GOOGLE_EMAIL_ADDRESS"`
	GoogleEmailSecret         string   `env:"GOOGLE_EMAIL_SECRET"`
	PasswordResetRedirectPage string   `env:"PASSWORD_RESET_REDIRECT_PAGE"`
	AllowedOrigins            []string `env:"ALLOWED_ORIGINS" envSeparator:","`
	JWTSecret                 string   `env:"JWT_SECRET"`
}

func InitConfig() *Config {
	var config Config
	if err := env.Parse(&config); err != nil {
		log.Fatalf("env.Parse() in config failed. Error:'%v'", err)
	}

	config.AllowedOrigins = []string{"*"}
	return &config
}
