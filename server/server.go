package server

import (
	"fmt"
	"net/http"

	"github.com/miguelgz36/MamuroSearchService/server/routes"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func InitServer() {
	port := ":8080"

	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	r.Use(middleware.Logger)
	r.Use(cors.Handler)

	routes.InitRoutes(r)

	error := http.ListenAndServe(port, r)
	if error != nil {
		panic(error)
	}
	fmt.Println("server is running at port: " + port)
}
