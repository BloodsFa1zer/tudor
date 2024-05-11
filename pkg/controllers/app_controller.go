package controllers

import middleware "study_marketplace/pkg/middlewares"

type AppController struct {
	UserControllerInterface
	AdvertisementsControllerInterface
	AuthControllerInterface
	CategoriesControllerInterface

	middleware.Middleware
}
