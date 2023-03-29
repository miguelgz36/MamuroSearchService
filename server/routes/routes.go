package routes

import (
	"github.com/go-chi/chi"
	"github.com/miguelgz36/MamuroSearchService/server/controllers"
)

func InitRoutes(r *chi.Mux) {
	r.Get("/search", controllers.Search)
}
