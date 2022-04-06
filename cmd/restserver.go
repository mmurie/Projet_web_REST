package main

import (
	"fmt"
	"internal/entities"
	"internal/handlers"
	"internal/persistence"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println(" --- PROJET WEB REST --- ")

	r := mux.NewRouter()

	leStudentDAO := new(persistence.StudentDAOMemory)
	leLanguageDAO := new(persistence.LanguageDAOMemory)

	student1 := entities.NewStudent(1, "Michel", "Baie", 25, "js")
	student2 := entities.NewStudent(2, "Paul", "Patine", 19, "c")
	students := make(map[int]entities.Student)
	students[1] = student1
	students[2] = student2
	studentDAO := new(persistence.StudentDAOMemory)
	studentDAO.Students = students

	language1 := entities.NewLanguage("js", "JavaScript")
	language2 := entities.NewLanguage("c", "C")
	languages := make(map[string]entities.Language)
	languages["js"] = language1
	languages["c"] = language2
	languageDAO := new(persistence.LanguageDAOMemory)
	languageDAO.Languages = languages

	handlers.InitializeStudentsRoutes(r, leStudentDAO)
	handlers.InitializeLanguagesRoutes(r, leLanguageDAO)

	err := http.ListenAndServe(":8000", r)

	if err != nil {
		log.Fatal(err)
	}
}
