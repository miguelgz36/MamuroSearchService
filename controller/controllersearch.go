package controller

import (
	"net/http"
)

func Search(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
