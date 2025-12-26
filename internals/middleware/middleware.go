package middleware

import (
	"crud_api/configs"
	"fmt"
	"net/http"
	"time"
)

func SecLayer(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		apikey := r.Header.Get("X-API-KEY")
		if apikey != configs.Key{
			http.Error(w, "Incorrect API Key", http.StatusUnauthorized)
			return 
		}
		fmt.Printf("Started %s %s\n",r.Method, r.URL.Path)

		next.ServeHTTP(w, r)

		fmt.Printf("Completed in %v\n", time.Since(start))
	})
}