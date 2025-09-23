package product

import (
	"ecommerce/rest/middleware"
)

type Handler struct{
	middlewares *middleware.Middlewares
	svc Service
}

func NewHandler(
	middlewares *middleware.Middlewares,
	svc Service,
	) *Handler {
	return &Handler{
		middlewares: middlewares,
		svc: svc,
	}
}