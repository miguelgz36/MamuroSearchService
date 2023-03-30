package controllers

import (
	"net/http"

	"github.com/miguelgz36/MamuroSearchService/server/services"
)

func Search(w http.ResponseWriter, r *http.Request) {
	w.Write(services.Search())
}
