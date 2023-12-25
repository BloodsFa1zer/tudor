package registry

import (
	"study_marketplace/database/queries"
	"study_marketplace/pkg/infrastructure/database"
)

func repoRegistry(dburl string) *queries.Queries {
	conn := database.ConnectDataBase(dburl)
	db := queries.New(conn)
	return db
}
