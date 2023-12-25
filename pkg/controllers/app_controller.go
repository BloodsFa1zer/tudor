package controllers

import middleware "study_marketplace/pkg/middlewares"

type AppController struct {
	UserController
	AdvertisementsController
	AuthController
	CategoriesController

	middleware.Middleware
}
