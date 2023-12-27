package registry

import (
	"study_marketplace/database/queries"
	"study_marketplace/pkg/controllers"
	"study_marketplace/pkg/domen/models"
	config "study_marketplace/pkg/infrastructure/config"
	middleware "study_marketplace/pkg/middlewares"
	"study_marketplace/pkg/repositories"
	"study_marketplace/pkg/services"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type registry struct {
	queries *queries.Queries
	config  *config.Config
}

type Registry interface {
	NewAppController() *controllers.AppController
}

func NewRegistry(config *config.Config, db *pgx.Conn) Registry {
	return &registry{
		queries: queries.New(db),
		config:  config,
	}
}

func (r *registry) NewAppController() *controllers.AppController {
	return &controllers.AppController{
		UserController:           r.userRegister(),
		AdvertisementsController: r.advRegister(),
		AuthController:           r.authRegister(),
		CategoriesController:     r.categoriesRegister(),

		Middleware: r.middlewareRegister(),
	}
}

func (r *registry) userRegister() controllers.UserController {
	return controllers.NewUsersController(
		services.NewUserService(r.config, r.genTokFunc(), r.hashPasFunc(), r.comparePasFunc(),
			repositories.NewUsersRepository(r.queries)))
}

func (r *registry) categoriesRegister() controllers.CategoriesController {
	return controllers.NewCatController(
		services.NewCategoriesService(
			repositories.NewCategoriesRepository(r.queries)))
}

func (r *registry) authRegister() controllers.AuthController {
	return controllers.NewAuthController(r.config.RedirectUrl,
		services.NewUserService(r.config, r.genTokFunc(), r.hashPasFunc(), r.comparePasFunc(),
			repositories.NewUsersRepository(r.queries)))
}

func (r *registry) advRegister() controllers.AdvertisementsController {
	return controllers.NewAdvertisementsController(
		services.NewAdvertisementService(
			repositories.NewAdvertisementsRepository(r.queries)))
}

// middlewareRegister returns a middleware.Middleware that is created using the configuration stored in the registry.
func (r *registry) middlewareRegister() middleware.Middleware {
	return middleware.NewMiddleware(r.config)
}

func (r *registry) genTokFunc() func(userid int64, email string) (string, error) {
	return func(userid int64, userName string) (string, error) {
		claims := &models.AuthClaims{
			UserID: userid,
			Email:  userName,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 10)), // Set token expiration time
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte(r.config.JWTSecret))
		if err != nil {
			return "", err
		}
		return tokenString, nil
	}
}

func (r *registry) hashPasFunc() func(password string) string {
	return func(password string) string {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		return string(hashedPassword)
	}
}

func (r *registry) comparePasFunc() func(hashedPassword string, candidatePassword string) error {
	return func(hashedPassword string, candidatePassword string) error {
		return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
	}
}
