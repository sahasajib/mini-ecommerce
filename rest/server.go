package rest

import (
	"ecommerce/config"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middleware"
	"fmt"
	"net/http"
	"os"
	"strconv"
)


type Server struct {
	cnf *config.Config
	productHandler *product.Handler
	userHandler *user.Handler
}

func NewServer(
	cnf *config.Config,
	productHandler *product.Handler, 
	userHandler *user.Handler,
	) *Server {
	return &Server{
		cnf: cnf,
		productHandler: productHandler,
		userHandler: userHandler,
	}
}

func (s *Server) Start(){
	mux := http.NewServeMux()
	
	manager := middleware.NewManager()
	manager.Use(
		middleware.Cors,
		middleware.Preflight,
		middleware.Logger,
	)
	

	wrappedMux := manager.WrapMux(mux)

	//Routes(mux, manager)
	s.productHandler.RegisterRoutes(mux, manager)
	s.userHandler.RegisterRoutes(mux, manager)

	addr := ":" + strconv.Itoa(s.cnf.HttpPort)
	fmt.Println("Server running on: ", addr)
	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil{
		fmt.Println("Error starting the server.......", err)
		os.Exit(1)
	}
}