package main

import (
	"backend/routes"
	"bytes"
	"io"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/run", withCORS(routes.RunHandler))

	logged := logMiddleware(mux)

	log.Println("Listening to http://localhost:8001")
	log.Fatal(http.ListenAndServe(":8001", logged))
}

func withCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			return
		}
		next(w, r)
	}
}

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll((r.Body))
		r.Body = io.NopCloser(io.MultiReader(bytes.NewReader(body)))

		log.Println(">>>", r.Method, r.URL.Path)
		log.Println("Payload", string(body))

		next.ServeHTTP(w, r)
	})
}
