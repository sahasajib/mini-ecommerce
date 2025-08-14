package cmd

import (
	"ecommerce/handlers"
	"ecommerce/middleware"
	"net/http"
)

func Routes(mux *http.ServeMux, manager *middleware.Manager){
	mux.Handle("GET /products",  manager.With(http.HandlerFunc(handlers.GetProducts)))
	mux.Handle("POST /products",  manager.With(http.HandlerFunc(handlers.CreteProducts)))
	mux.Handle("GET /products/{id}",  manager.With(http.HandlerFunc(handlers.GetProductByID)))
}