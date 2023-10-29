package application

import (
	"github.com/gorilla/mux"
)

func RegisterApiRouters(r *mux.Router) {
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/products", productsHandler)
	r.HandleFunc("/articles", articlesHandler)
}
