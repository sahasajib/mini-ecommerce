package db

import (
	"ecommerce/config"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString(cnf *config.DBConfig) string {
	connString := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s",
	cnf.User,
	cnf.Password,
	cnf.Host,
	cnf.Port,
	cnf.Name,
	) 
	if !cnf.EnableSSLMode{
		connString += " sslmode=disable"
	}
	return connString
}

func NewConnection(cnf *config.DBConfig) (*sqlx.DB, error){
	dbSource := GetConnectionString(cnf)
	log.Println("DB Source: ", dbSource)
	dbCon, err := sqlx.Connect("postgres", dbSource)
	log.Println("DB Connection: ", dbCon)
	if err != nil{
		fmt.Println("Failed to connect to db", err)
		return nil, err
	}
	return dbCon, nil
}