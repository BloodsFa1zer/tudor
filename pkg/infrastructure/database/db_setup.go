package database

import (
	"context"
	"log"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDataBase(dburl string) *pgxpool.Pool {
	dbpool, err := pgxpool.New(context.Background(), dburl)
	if err != nil {
		log.Fatal("create database conection error:" + err.Error())
	}

	return dbpool
}
