package routers

import (
	config "study_marketplace/config"
	"study_marketplace/controllers"
	_ "study_marketplace/docs"
	middleware "study_marketplace/middlewares"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(conf *config.Config, server *gin.Engine, ac *controllers.AppController) {

	server.Use(middleware.CORS(conf))

	docs_url := ginSwagger.URL(conf.DocsHostname + "/api/docs/doc.json")

	api := server.Group("/api")

	api.GET("/", controllers.HealthCheck)

	api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, docs_url))
	api.POST("/auth/register", ac.UserRegister)
	api.POST("/auth/login", ac.UserLogin)
	api.GET("/auth/login-google", ac.LoginGoogle)
	api.GET("/auth/login-google-callback", ac.LoginGoogleCallback)
	// api.GET("/auth/login-facebook", a.LoginFacebook)
	api.POST("/auth/reset-password", ac.PasswordReset)

	protected := server.Group("/protected")

	protected.Use(middleware.AuthMiddleware())
	protected.GET("/userinfo", ac.UserInfo)
	protected.PATCH("/create-password", ac.PasswordCreate)

	// categories block
	categories := server.Group("/open/categories")
	categories.GET("/getall", ac.CatGetAll)

	// advertisements block
	// open advertisements endpoints
	advertisements := server.Group("/open/advertisements")
	advertisements.GET("/getall", ac.AdvGetAll)
	advertisements.GET("/getbyid/:id", ac.AdvGetByID)

	// protected advertisements endpoints
	protected.POST("/advertisement-create", ac.AdvCreate)
	protected.PATCH("/advertisement-patch", ac.AdvPatch)
	protected.DELETE("/advertisement-delete", ac.AdvDelete)
	protected.POST("/advertisement-filter", ac.AdvGetFiltered)
	protected.GET("/advertisement-getmy", ac.AdvGetMy)

	protected.Use(middleware.PasswordMiddleware(ac))
	protected.PATCH("/user-patch", ac.UserPatch)

}
