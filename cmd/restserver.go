package main

import (
	"fmt"
	"internal/entities"
	"internal/handlers"
	"log"
	"net/http"
)

func main() {
	fmt.Println(" --- PROJET WEB REST --- ")

	language1 := entities.NewLanguage("js", "JavaScript")
	language2 := entities.NewLanguage("c", "C")

	fmt.Println(entities.Language.String(language1))
	fmt.Println(entities.Language.String(language2))

	err := http.ListenAndServe(":8000", nil)

	handlers.InitializeLanguageRoutes()

	if err != nil {
		log.Fatal(err)
	}
}
