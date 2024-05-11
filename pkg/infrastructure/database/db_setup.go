package database

import (
	"context"
	"log"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4"
)

func ConnectDataBase(dbUrl string) *pgx.Conn {
	db, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		log.Fatal("create database conection error:" + err.Error())
	}
	return db
}
