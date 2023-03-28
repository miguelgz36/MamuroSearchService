package routes

import (
	"github.com/miguelgz36/MamuroSearchService/controller"

	"github.com/go-chi/chi"
)

func InitRoutes(r *chi.Mux) {
	r.Get("/search", controller.Search)
}
