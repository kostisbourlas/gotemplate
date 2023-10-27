package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kostisbourlas/gotemplate/application"
)

func main() {
	r := mux.NewRouter()
	application.RegisterRouters(r)

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
