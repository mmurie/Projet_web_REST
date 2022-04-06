package handlers

import (
	"encoding/json"
	"fmt"
	"internal/hardcodedData"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" - - - - - - - - - - - ")
	fmt.Println("func GetAllStudents")

	jsonString, err := json.Marshal(hardcodedData.GetAllDatas())
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, string(jsonString))

}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" - - - - - - - - - - - ")
	fmt.Println("func GetStudent")

	vars := mux.Vars(r)
	code := vars["code"]

	jsonString, err := json.Marshal(hardcodedData.GetData(code))
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

func InitializeStudentsRoutes(r *mux.Router) {
	r.HandleFunc("/rest/students/{id}", GetStudent).Methods("GET")
	r.HandleFunc("/rest/students", GetAllStudents).Methods("GET")
	r.HandleFunc("/rest/students", AddStudent).Methods("POST")
	r.HandleFunc("/rest/students", EditStudent).Methods("PUT")
	r.HandleFunc("/rest/students/{id}", DeleteStudent).Methods("DELETE")

}
