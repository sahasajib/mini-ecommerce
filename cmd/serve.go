package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/repo"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middleware"
	"fmt"
	"os"
)

func Serve() {
	cnf := config.GetConfig()

	dbCon, err := db.NewConnection()
	if err != nil{
		fmt.Println("Failed to connect to db")
		os.Exit(1)
	}

	middlewares := middleware.NewMiddlewares(cnf)

	productRepo := repo.NewProductRepo(dbCon)
	productHandler := product.NewHandler(middlewares, productRepo)

	userRepo := repo.NewUserRepo(dbCon)
	userHandler := user.NewHandler(cnf, userRepo)
	
	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
	)
	server.Start()
}