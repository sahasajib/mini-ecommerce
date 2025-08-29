package rest

import (
	"ecommerce/config"
	"ecommerce/rest/middleware"
	"fmt"
	"net/http"
	"os"
	"strconv"
)


func Start(cnf config.Config){
	mux := http.NewServeMux()
	
	manager := middleware.NewManager()
	manager.Use(
		middleware.Cors,
		middleware.Preflight,
		middleware.Logger,
	)
	

	wrappedMux := manager.WrapMux(mux)

	Routes(mux, manager)
	

	addr := ":" + strconv.Itoa(cnf.HttpPort)
	fmt.Println("Server running on: ", addr)
	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil{
		fmt.Println("Error starting the server.......", err)
		os.Exit(1)
	}
}