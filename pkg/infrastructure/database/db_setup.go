package database

import (
	"context"
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDataBase(dburl string) *pgxpool.Pool {
	dbpool, err := pgxpool.New(context.Background(), dburl)
	if err != nil {
		log.Fatal("create database conection error:" + err.Error())
	}
	migration(dburl)
	return dbpool
}

func migration(dburl string) {
	db, err := sql.Open("postgres", dburl)
	if err != nil {
		log.Fatal("open database migration error:" + err.Error())
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal("config database migration error:" + err.Error())
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file:///database/migrations",
		"postgres", driver)
	if err != nil {
		log.Fatal("database migration error:" + err.Error())
	}
	m.Up()
}
