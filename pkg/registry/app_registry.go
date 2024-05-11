package registry

import (
	"study_marketplace/pkg/controllers"
	middleware "study_marketplace/pkg/middlewares"
	"study_marketplace/pkg/repositories"
	"study_marketplace/pkg/services"
)

func userRegister(r *registry) controllers.UserController {
	return controllers.NewUsersController(
		services.NewUserService(
			generateTokenFunc(r.config.JWTSecret),
			hashPasswordFunc(),
			compareHashedPasswordFunc(),
			emailSenderFunc(r.config.RedirectUrl, r.config.GoogleEmailSenderName, r.config.GoogleEmailAddress, r.config.GoogleEmailSecret),
			repositories.NewUsersRepository(r.queries)), r.config.BasicAppUrl)
}

func categoriesRegister(r *registry) controllers.CategoriesController {
	return controllers.NewCatController(
		services.NewCategoriesService(
			repositories.NewCategoriesRepository(r.queries)))
}

func authRegister(r *registry) controllers.AuthController {
	return controllers.NewAuthController(r.config.RedirectUrl, callBackFunc(),
		services.NewAuthService(
			generateTokenFunc(r.config.JWTSecret),
			repositories.NewAuthRepository(r.queries)))
}

func advRegister(r *registry) controllers.AdvertisementsController {
	return controllers.NewAdvertisementsController(
		services.NewAdvertisementService(
			repositories.NewAdvertisementsRepository(r.queries)))
}

// middlewareRegister returns a middleware.Middleware that is created using the configuration stored in the registry.
func middlewareRegister(r *registry) middleware.Middleware {
	return middleware.NewMiddleware(r.config)
}
