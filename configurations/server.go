package configurations

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func InitServer() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	//routes.InitRoutes()

	http.ListenAndServe(":8080", r)
}
