package controllers

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"study_marketplace/domen/models"
	"study_marketplace/services"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type AuthController interface {
	LoginGoogle(ctx *gin.Context)
	LoginGoogleCallback(ctx *gin.Context)
}

const ProtocolPrefix = "https"
const GoogleCallbackUrl = "/api/auth/login-google-callback"
const googleCookieName = "oauthstate"
const GoogleQueryNameState = "state"
const GoogleQueryNameCode = "code"
const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

type authController struct {
	redirectPage      string
	googleOauthConfig *oauth2.Config
	userService       services.UserService
}

func NewAuthController(redirectPage string, us services.UserService) AuthController {

	url := ProtocolPrefix + "://" + os.Getenv("GOOGLE_CALLBACK_DOMAIN") + GoogleCallbackUrl

	return &authController{
		redirectPage: redirectPage,
		googleOauthConfig: &oauth2.Config{
			RedirectURL:  url,
			ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
			Scopes: []string{
				"https://www.googleapis.com/auth/userinfo.email",
				"https://www.googleapis.com/auth/userinfo.profile",
			},
			Endpoint: google.Endpoint,
		},
		userService: us,
	}
}

// b := make([]byte, 16)
// rand.Read(b)
// googleCookieValue := base64.URLEncoding.EncodeToString(b)

func (t *authController) LoginGoogle(ctx *gin.Context) {
	b := make([]byte, 16)
	rand.Read(b)
	googleCookieValue := base64.URLEncoding.EncodeToString(b)
	maxAge := 3600
	path := ""
	domain := os.Getenv("GOOGLE_CALLBACK_DOMAIN")
	secure := false
	httpOnly := true
	ctx.SetCookie(googleCookieName, googleCookieValue, maxAge, path, domain, secure, httpOnly)
	authURL := t.googleOauthConfig.AuthCodeURL(googleCookieValue)
	ctx.Redirect(http.StatusTemporaryRedirect, authURL)
}

func (t *authController) LoginGoogleCallback(ctx *gin.Context) {
	googleCookieValue, _ := ctx.Cookie(googleCookieName)
	googleQueryNameState := ctx.Query(GoogleQueryNameState)
	if googleCookieValue != googleQueryNameState {
		log.Println("WARNING: invalid oauth google state")
	}

	codeStr := ctx.Query(GoogleQueryNameCode)
	token, err := t.googleOauthConfig.Exchange(context.Background(), codeStr)
	if err != nil {
		fmt.Println("code exchange wrong", err.Error())
	}

	response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		fmt.Println("failed getting user info", err.Error())
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("failed read user info", err.Error())
	}

	var userInfo models.GoogleResponse
	err = json.Unmarshal(body, &userInfo)
	if err != nil {
		fmt.Println("failed parse response body")
	}
	// user, err := t.userService.GetOrCreateUser(ctx, userInfo)

	// if err != nil {
	// 	fmt.Println("Failed to get user by email.")
	// }

	// tokenJWT, err := t.userService.GenToken(user.ID, user.Name)
	// if err != nil {
	// 	fmt.Println("Failed to generate token")
	// }

	q := url.Values{}
	q.Set("token", "string(tokenJWT)")
	location := url.URL{Path: t.redirectPage, RawQuery: q.Encode()}
	ctx.Redirect(http.StatusPermanentRedirect, location.RequestURI())
}
