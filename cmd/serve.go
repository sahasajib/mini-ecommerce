package cmd

import (
	"ecommerce/global_routes"
	"ecommerce/handlers"
	"fmt"
	"net/http"
)


func Serve() {
	mux := http.NewServeMux()

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /products", http.HandlerFunc(handlers.CreteProducts))
	mux.Handle("GET /products/{id}", http.HandlerFunc(handlers.GetProductByID))
	
	globalRout := global_routes.GlobalRouter(mux)
	fmt.Println("Server running on: 8080")

	err := http.ListenAndServe(":8080", globalRout)
	if err != nil{
		fmt.Println("Error starting the server.......", err)
	}
}