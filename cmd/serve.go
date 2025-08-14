package cmd

import (
	"ecommerce/global_routes"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)


func Serve() {
	manager := middleware.NewManager()
	manager.Use(middleware.Logger)
	mux := http.NewServeMux()


	Routes(mux, manager)
	
	
	globalRout := global_routes.GlobalRouter(mux)
	fmt.Println("Server running on: 8080")

	err := http.ListenAndServe(":8080", globalRout)
	if err != nil{
		fmt.Println("Error starting the server.......", err)
	}
}