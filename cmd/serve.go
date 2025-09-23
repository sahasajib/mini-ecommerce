package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/product"
	"ecommerce/repo"
	"ecommerce/rest"
	productHandler "ecommerce/rest/handlers/product"
	usrHandler "ecommerce/rest/handlers/user"
	"ecommerce/rest/middleware"
	"ecommerce/user"
	"fmt"
	"os"
)

func Serve() {
	cnf := config.GetConfig()
	//fmt.Println("%+v", cnf.DB)

	dbCon, err := db.NewConnection(cnf.DB)
	if err != nil{
		fmt.Println("Failed to connect to db")
		os.Exit(1)
	}
	err = db.MigrateDB(dbCon, "./migrations")
	if err != nil{
		fmt.Println("DB Migration failed", err)
		os.Exit(1)
	}
	middlewares := middleware.NewMiddlewares(cnf)

	//repos
	productRepo := repo.NewProductRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)

	//domain
	usrSvc := user.NewService(userRepo)
	prdctSvc := product.NewService(productRepo)

	productHandler := productHandler.NewHandler(middlewares, prdctSvc)

	
	userHandler := usrHandler.NewHandler(cnf, usrSvc)
	
	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
	)
	server.Start()
}