package handlers

import (
	"encoding/json"
	"fmt"
	"internal/persistence"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type StudentsHandlers struct {
	DAO persistence.StudentDAO
}

func (sh StudentsHandlers) GetAllStudents(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" - - - - - - - - - - - ")
	fmt.Println("func GetAllStudents")

	jsonString, err := json.Marshal(sh.DAO.FindAll())
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, string(jsonString))
}

func (sh StudentsHandlers) GetStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" - - - - - - - - - - - ")
	fmt.Println("func GetStudent")

	vars := mux.Vars(r)
	id := vars["id"]

	id_int, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
	}

	jsonString, err := json.Marshal(sh.DAO.Find(id_int))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, string(jsonString))
}

func AddStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" - - - - - - - - - - - ")
	fmt.Println("func AddStudent")
}

func EditStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" - - - - - - - - - - - ")
	fmt.Println("func EditStudent")
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" - - - - - - - - - - - ")
	fmt.Println("func DeleteStudent")
}

func (sh StudentsHandlers) InitializeStudentsRoutes(r *mux.Router, studentDAO persistence.StudentDAO) {
	r.HandleFunc("/rest/students/{id}", sh.GetStudent).Methods("GET")
	r.HandleFunc("/rest/students", sh.GetAllStudents).Methods("GET")
	r.HandleFunc("/rest/students", AddStudent).Methods("POST")
	r.HandleFunc("/rest/students", EditStudent).Methods("PUT")
	r.HandleFunc("/rest/students/{id}", DeleteStudent).Methods("DELETE")

}
