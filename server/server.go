package server

import (
	"net/http"

	"github.com/miguelgz36/MamuroSearchService/server/routes"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func InitServer() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	routes.InitRoutes(r)

	error := http.ListenAndServe(":8081", r)
	if error != nil {
		panic(error)
	}
}
