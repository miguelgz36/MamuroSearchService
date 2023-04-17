package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/miguelgz36/MamuroSearchService/server/services"
)

func Search(w http.ResponseWriter, r *http.Request) {
	textToSearch := r.URL.Query().Get("text")
	if textToSearch == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Text to search missing"))
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Page missing or badformed"))
	}
	fmt.Println("Request search: " + textToSearch)

	result, err := services.Search(textToSearch, page)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(result)
	}

	w.Write(result)
	w.WriteHeader(http.StatusOK)
}
