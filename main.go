package main

import (
	"ecommerce/cmd"
)

// func crosHandler(next http.Handler) http.Handler {
// 	handleCors := func (w http.ResponseWriter, r *http.Request){
// 		next.ServeHTTP(w, r)
// 	}
// 	handler := http.HandlerFunc(handleCors)
// 	return handler
// }

func main() {
	cmd.Serve()
}
