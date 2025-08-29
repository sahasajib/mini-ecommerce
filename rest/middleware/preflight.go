package middleware

import "net/http"


func Preflight(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
		// Preflight
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		} 

		next.ServeHTTP(w, r)
	})
}