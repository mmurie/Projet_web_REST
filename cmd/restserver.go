package main

import (
	"fmt"
	"internal/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println(" --- PROJET WEB REST --- ")

	r := mux.NewRouter()

	handlers.InitializeLanguageRoutes(r)

	err := http.ListenAndServe(":8000", r)

	if err != nil {
		log.Fatal(err)
	}
}
