package configurations

import (
	"net/http"

	"github.com/miguelgz36/MamuroSearchService/configurations/routes"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func InitServer() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	routes.InitRoutes(r)

	http.ListenAndServe(":8080", r)
}
