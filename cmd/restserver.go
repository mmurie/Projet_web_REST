// PROJET API WEB REST EN GO
//
// Projet de LP MiAR de création d'une API REST en Go
//
// Terms Of Service:
//
//		Schemes: http, https
//		Host: localhost:8000
//		BasePath: /rest
//		Version: 1.0.0
//
//		Consumes:
//		- application/json
//
//		Produces:
//		- application/json
//
// swagger:meta
package main

import (
	"encoding/json"
	"fmt"
	"internal/entities"
	"internal/handlers"
	"internal/persistence"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println(" --- PROJET WEB REST --- ")

	r := mux.NewRouter()

	//Version avec DAOMemory
	/*
		leStudentDAO := new(persistence.StudentDAOMemory)
		leLanguageDAO := new(persistence.LanguageDAOMemory)

		student1 := entities.NewStudent(1, "Michel", "Baie", 25, "js")
		student2 := entities.NewStudent(2, "Paul", "Patine", 19, "c")
		students := make(map[int]entities.Student)
		students[1] = student1
		students[2] = student2
		leStudentDAO.Students = students

		language1 := entities.NewLanguage("js", "JavaScript")
		language2 := entities.NewLanguage("c", "C")
		languages := make(map[string]entities.Language)
		languages["js"] = language1
		languages["c"] = language2
		leLanguageDAO.Languages = languages
	*/

	bdd := new(persistence.BoltDb)
	bdd.DbOpen("bdd.db")
	defer bdd.DbClose()
	bdd.DbCreateBucket("languages")
	bdd.DbCreateBucket("students")

	var encoded []byte
	var encoded_err error

	//Students
	leStudentDAO := new(persistence.StudentDAOBolt)

	student1 := entities.NewStudent(1, "Michel", "Baie", 25, "js")
	student2 := entities.NewStudent(2, "Paul", "Patine", 19, "c")

	encoded, encoded_err = json.Marshal(student1)
	if encoded_err != nil {
		log.Fatal(encoded_err)
	}
	bdd.DbPut("students", strconv.Itoa(student1.Id), string(encoded))

	encoded, encoded_err = json.Marshal(student2)
	if encoded_err != nil {
		log.Fatal(encoded_err)
	}
	bdd.DbPut("students", strconv.Itoa(student2.Id), string(encoded))

	leStudentDAO.Dbms = *bdd

	//Languages
	leLanguageDAO := new(persistence.LanguageDAOBolt)

	language1 := entities.NewLanguage("js", "JavaScript")
	language2 := entities.NewLanguage("c", "C")

	encoded, encoded_err = json.Marshal(language1)
	if encoded_err != nil {
		log.Fatal(encoded_err)
	}
	bdd.DbPut("languages", language1.Code, string(encoded))

	encoded, encoded_err = json.Marshal(language2)
	if encoded_err != nil {
		log.Fatal(encoded_err)
	}
	bdd.DbPut("languages", language2.Code, string(encoded))

	leLanguageDAO.Dbms = *bdd

	//Handlers
	studentsHandlers := new(handlers.StudentsHandlers)
	studentsHandlers.DAO = leStudentDAO
	studentsHandlers.InitializeStudentsRoutes(r)
	languageHandlers := new(handlers.LanguagesHandlers)
	languageHandlers.DAO = leLanguageDAO
	languageHandlers.InitializeLanguagesRoutes(r)

	//Route pour la documentation
	fs := http.FileServer(http.Dir("./swagger"))
	r.PathPrefix("/docs/").Handler(http.StripPrefix("/docs/", fs))

	err := http.ListenAndServe(":8000", r)

	if err != nil {
		log.Fatal(err)
	}
}
