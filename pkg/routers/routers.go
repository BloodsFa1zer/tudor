package routers

import (
	_ "study_marketplace/docs"
	"study_marketplace/pkg/controllers"
	config "study_marketplace/pkg/infrastructure/config"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(conf *config.Config, server *gin.Engine, a *controllers.AppController) {
	server.Use(a.CORS())
	api := server.Group("/api")

	api.GET("/", controllers.HealthCheck)
	api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL(conf.ServerHostname+"/api/docs/doc.json")))
	api.POST("/auth/register", a.UserRegister)
	api.POST("/auth/login", a.UserLogin)
	api.GET("/auth/:provider", a.AuthWithProvider)
	api.GET("/auth/:provider/callback", a.AuthWithProviderCallback)
	api.POST("/auth/reset-password", a.PasswordReset)

	protected := server.Group("/protected")

	protected.Use(a.AuthMiddleware())
	protected.GET("/userinfo", a.UserInfo)
	protected.PATCH("/create-password", a.PasswordCreate)

	// categories block
	server.GET("open/allcategories", a.CatGetAll)

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
