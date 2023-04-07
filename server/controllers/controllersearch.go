package controllers

import (
	"fmt"
	"net/http"

	"github.com/miguelgz36/MamuroSearchService/server/services"
)

func Search(w http.ResponseWriter, r *http.Request) {
	textToSearch := r.URL.Query().Get("text")
	fmt.Println("TEXT: " + textToSearch)
	w.Write(services.Search(textToSearch))
}
