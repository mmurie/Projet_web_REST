package handlers

import (
	"encoding/json"
	"fmt"
	"internal/entities"
	"internal/persistence"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type StudentsHandlers struct {
	DAO persistence.StudentDAO
}

// swagger:operation GET /students students GetAllStudents
// ---
// summary: Renvoie la liste de tous les étudiants
// responses:
//   "200":
//     "$ref": "#/responses/studentsRes"
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

func (sh StudentsHandlers) AddStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" - - - - - - - - - - - ")
	fmt.Println("func AddStudent")

	student := entities.NewStudent(0, "", "", 0, "")
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &student)

	isOk, err := json.Marshal(sh.DAO.Create(student))

	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, string(isOk))
}

func (sh StudentsHandlers) EditStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" - - - - - - - - - - - ")
	fmt.Println("func EditStudent")

	student := entities.NewStudent(0, "", "", 0, "")
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &student)

	isOk, err := json.Marshal(sh.DAO.Update(student))

	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, string(isOk))
}

func (sh StudentsHandlers) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" - - - - - - - - - - - ")
	fmt.Println("func DeleteStudent")

	vars := mux.Vars(r)
	id := vars["id"]

	id_int, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
	}

	isOk, err := json.Marshal(sh.DAO.Delete(id_int))

	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, string(isOk))

}

func (sh StudentsHandlers) InitializeStudentsRoutes(r *mux.Router) {
	r.HandleFunc("/rest/students/{id}", sh.GetStudent).Methods("GET")
	r.HandleFunc("/rest/students", sh.GetAllStudents).Methods("GET")
	r.HandleFunc("/rest/students", sh.AddStudent).Methods("POST")
	r.HandleFunc("/rest/students", sh.EditStudent).Methods("PUT")
	r.HandleFunc("/rest/students/{id}", sh.DeleteStudent).Methods("DELETE")
}
