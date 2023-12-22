package routers

import (
	"study_marketplace/controllers"
	_ "study_marketplace/docs"
	config "study_marketplace/internal/infrastructure/config"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(conf *config.Config, server *gin.Engine, a *controllers.AppController) {

	server.Use(a.CORS())

	docs_url := ginSwagger.URL(conf.DocsHostname + "/api/docs/doc.json")

	api := server.Group("/api")

	api.GET("/", controllers.HealthCheck)

	api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, docs_url))
	api.POST("/auth/register", a.UserRegister)
	api.POST("/auth/login", a.UserLogin)
	api.GET("/auth/login-google", a.LoginGoogle)
	api.GET("/auth/login-google-callback", a.LoginGoogleCallback)
	// api.GET("/auth/login-facebook", a.LoginFacebook)
	api.POST("/auth/reset-password", a.PasswordReset)

	protected := server.Group("/protected")

	protected.Use(a.AuthMiddleware())
	protected.GET("/userinfo", a.UserInfo)
	protected.PATCH("/create-password", a.PasswordCreate)

	// categories block
	categories := server.Group("/open/categories")
	categories.GET("/getall", a.CatGetAll)

	// advertisements block
	// open advertisements endpoints
	advertisements := server.Group("/open/advertisements")
	advertisements.GET("/getall", a.AdvGetAll)
	advertisements.GET("/getbyid/:id", a.AdvGetByID)

	// protected advertisements endpoints
	protected.POST("/advertisement-create", a.AdvCreate)
	protected.PATCH("/advertisement-patch", a.AdvPatch)
	protected.DELETE("/advertisement-delete", a.AdvDelete)
	protected.POST("/advertisement-filter", a.AdvGetFiltered)
	protected.GET("/advertisement-getmy", a.AdvGetMy)

	protected.Use(a.PasswordMiddleware())
	protected.PATCH("/user-patch", a.UserPatch)

}
