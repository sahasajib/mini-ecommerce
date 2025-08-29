package rest

import (
	"ecommerce/rest/middleware"
	"ecommerce/rest/handlers"
	"net/http"
)

func Routes(mux *http.ServeMux, manager *middleware.Manager){
	mux.Handle("GET /products",  manager.With(http.HandlerFunc(handlers.GetProducts)))
	mux.Handle("POST /products",  manager.With(http.HandlerFunc(handlers.CreteProducts)))
	mux.Handle("GET /products/{id}",  manager.With(http.HandlerFunc(handlers.GetProductByID)))
}