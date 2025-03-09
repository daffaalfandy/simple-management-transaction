package main

import (
	app "backend/cmd"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	r := app.Init()

	fmt.Println("Server running on port " + os.Getenv("APP_PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("APP_PORT"), r))
}
