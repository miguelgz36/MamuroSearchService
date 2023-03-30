package routes

import (
	"github.com/miguelgz36/MamuroSearchService/server/controllers"

	"github.com/go-chi/chi"
)

func InitRoutes(r *chi.Mux) {
	r.Get("/search", controllers.Search)
}
