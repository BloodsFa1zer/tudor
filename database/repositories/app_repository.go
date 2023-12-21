package repositories

import (
	config "study_marketplace/config"
	"study_marketplace/internal/database/queries"
	"study_marketplace/internal/infrastructure/database"
)

func NewAppRepository(conf *config.Config) *queries.Queries {
	conn := database.ConnectDataBase(conf)
	db := queries.New(conn)
	return db
}
