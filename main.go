package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bishnupandey/golang-react-todo/router"
	"github.com/gorilla/mux"
)

func main() {
	r := router.Router()
	fmt.Println("Starting the server on port 9000...")
	mux.CORSMethodMiddleware(r)
	log.Fatal(http.ListenAndServe(":9000", r))
}
