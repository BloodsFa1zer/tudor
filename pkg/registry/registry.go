package registry

import (
	"study_marketplace/database/queries"
	"study_marketplace/pkg/controllers"
	config "study_marketplace/pkg/infrastructure/config"

	"github.com/jackc/pgx/v4"
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
		UserController:           userRegister(r),
		AdvertisementsController: advRegister(r),
		AuthController:           authRegister(r),
		CategoriesController:     categoriesRegister(r),

		Middleware: middlewareRegister(r),
	}
}
