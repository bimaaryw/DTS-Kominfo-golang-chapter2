package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "301100"
	dbname = "hactiv8-dts-go"
)

func StartDB() *gorm.DB {
	init := fmt.Sprintf("host=%s port=%d  user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(init))
	if err != nil {
		panic(err)
	}

	return db
}