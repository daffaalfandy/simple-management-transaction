package main

import (
	"fmt"
	"log"
	"net/http"

	"backend/config"
	"backend/internal/handler"

	"github.com/gorilla/mux"
)

func main() {
	config.ConnectDB()

	r := mux.NewRouter()
	r.HandleFunc("/users", handler.GetUsers).Methods("GET")

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
