package infrastructure

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kostisbourlas/gotemplate/application/api"
)

const SERVER_PORT = ":8080"

func runServer() {
	router := mux.NewRouter()
	application.RegisterApiRouters(router)

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(SERVER_PORT, router))
}
