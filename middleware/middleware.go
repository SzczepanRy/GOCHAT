package middleware

import (
	"fmt"
	"net/http"
)


func HeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	    w.Header().Set("Access-Control-Allow-Origin", "*")
        fmt.Println("headerWritten")
		next.ServeHTTP(w, r)
	})

}
