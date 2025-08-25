package cmd

import (
	
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	
	mux := http.NewServeMux()
	
	manager := middleware.NewManager()
	manager.Use(
		middleware.Cors,
		middleware.Preflight,
		middleware.Logger,
	)
	

	wrappedMux := manager.WrapMux(mux)

	Routes(mux, manager)
	fmt.Println("Server running on: 8080")

	err := http.ListenAndServe(":8080", wrappedMux)
	if err != nil{
		fmt.Println("Error starting the server.......", err)
	}
}