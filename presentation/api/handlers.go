package presentation

import (
	"encoding/json"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
  response := map[string]string{"Message": "The is no place like 127.0.0.1"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func catHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"Message": "This is a cat resource"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func dogHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"Message": "This is a dog resource"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
