package main

import (
	"backend/routes"
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/run", withCORS(routes.RunHandler))

	logged := logMiddleware(mux)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fs := http.FileServer(http.Dir("./frontend"))
	mux.HandleFunc("/", withCORS(func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join("./frontend", r.URL.Path)
		_, err := os.Stat(path)
		if err == nil && !strings.HasSuffix(r.URL.Path, "/") {
			fs.ServeHTTP(w, r)
			return
		}
		http.ServeFile(w, r, "./frontend/index.html")
	}))

	log.Println("Listening to http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, logged))
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
