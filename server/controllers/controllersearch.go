package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/miguelgz36/MamuroSearchService/server/services"
)

func Search(w http.ResponseWriter, r *http.Request) {
	textToSearch := r.URL.Query().Get("text")
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		panic(err)
	}
	fmt.Println("TEXT: " + textToSearch)
	w.Write(services.Search(textToSearch, page))
}
