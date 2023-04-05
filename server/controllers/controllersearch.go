package controllers

import (
	"net/http"

	"github.com/miguelgz36/MamuroSearchService/server/services"
)

func Search(w http.ResponseWriter, r *http.Request) {
	textToSearch := r.URL.Query().Get("text")
	w.Write(services.Search(textToSearch))
}
