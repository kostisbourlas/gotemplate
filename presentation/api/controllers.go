package presentation

import (
	"encoding/json"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"Message": "This is a cat!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"Message": "This is the products resource!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func articlesHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"Message": "This is the articles resource!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
