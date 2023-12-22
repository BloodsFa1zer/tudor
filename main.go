package main

import (
	"log"

	config "study_marketplace/internal/infrastructure/config"
	"study_marketplace/internal/infrastructure/server"
	"study_marketplace/internal/registry"
	"study_marketplace/routers"
)

// @title						Study marketplace API
// @version					0.0.1
// @description				Marketplace to connect students and teachers
// @termsOfService				[TODO]
// @contact.name				API Support
// @contact.url				[TODO]
// @contact.email				[TODO]
// @license.name				[TODO]
// @license.url				[TODO]
// @BasePath					/
// @schemes					http https
// @securityDefinitions.apiKey	JWT
// @in							header
// @name						Authorization
func main() {
	conf := config.InitConfig()

	s := server.NewServer()

	ac := registry.NewRegistry(conf).NewAppController()

	routers.SetupRouter(conf, s, ac)

	log.Fatal(s.Run(conf.ServerHostname))
}
