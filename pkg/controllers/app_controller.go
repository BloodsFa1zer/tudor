package controllers

import middleware "study_marketplace/middlewares"

type AppController struct {
	UserController
	AdvertisementsController
	AuthController
	CategoriesController

	middleware.Middleware
}
