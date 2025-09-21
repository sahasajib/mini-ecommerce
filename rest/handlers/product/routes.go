package product

import (
	"ecommerce/rest/middleware"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager){
	mux.Handle("GET /products",  manager.With(http.HandlerFunc(h.GetProducts)))
	mux.Handle("POST /products",  manager.With(http.HandlerFunc(h.CreteProduct), h.middlewares.AutenticateJWT))
	mux.Handle("GET /products/{id}",  manager.With(http.HandlerFunc(h.GetProduct)))
	mux.Handle("PUT /products/{id}",  manager.With(http.HandlerFunc(h.UpdateProduct), h.middlewares.AutenticateJWT))
	mux.Handle("DELETE /products/{id}",  manager.With(http.HandlerFunc(h.DeleteProduct), h.middlewares.AutenticateJWT))
}