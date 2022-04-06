package main

import (
	"fmt"
	"internal/entities"
	"internal/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println(" --- PROJET WEB REST --- ")

	language1 := entities.NewLanguage("js", "JavaScript")
	language2 := entities.NewLanguage("c", "C")

	fmt.Println(entities.Language.String(language1))
	fmt.Println(entities.Language.String(language2))

	r := mux.NewRouter()

	handlers.InitializeLanguageRoutes(r)

	err := http.ListenAndServe(":8000", r)

	if err != nil {
		log.Fatal(err)
	}
}
