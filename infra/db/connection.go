package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString() string {
	return "user=sajib password=12345678 host=localhost port=5432 dbname=ecom sslmode=disable"
}

func NewConnection() (*sqlx.DB, error){
	dbSource := GetConnectionString()
	log.Println("DB Source: ", dbSource)
	dbCon, err := sqlx.Connect("postgres", dbSource)
	log.Println("DB Connection: ", dbCon)
	if err != nil{
		fmt.Println("Failed to connect to db", err)
		return nil, err
	}
	return dbCon, nil
}