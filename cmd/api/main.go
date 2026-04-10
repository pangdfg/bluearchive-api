package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	srv := &http.Server{
		Addr:    ":" + "8080",
		Handler: r,
	}

	go func() {
		log.Println("Server started on port", "8080")
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

}