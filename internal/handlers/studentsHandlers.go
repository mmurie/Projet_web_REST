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

// swagger:operation GET /students Students GetAllStudents
// ---
// summary: Renvoie la liste de tous les étudiants
// responses:
//   "200":
//     "$ref": "#/responses/studentsRes"
//   "400":
//     "$ref": "#/responses/badReq"
func (sh StudentsHandlers) GetAllStudents(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" - - - - - - - - - - - ")
	fmt.Println("func GetAllStudents")

	jsonString, err := json.Marshal(sh.DAO.FindAll())
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, string(jsonString))
}

// swagger:operation GET /students/{id} Students GetStudents
// ---
// summary: Renvoie l'étudiant d'id spécifié
// parameters:
// - name: id
//   in: path
//   description: id de l'étudiant
//   type: int
//   required: true
// responses:
//   "200":
//     "$ref": "#/responses/studentRes"
//   "400":
//     "$ref": "#/responses/badReq"
//   "404":
//     "$ref": "#/responses/notFoundReq"
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

// swagger:operation POST /students Students AddStudent
// ---
// summary: Ajoute un nouvel étudiant
// parameters:
// - name: student
//   description: L'étudiant à ajouter
//   in: body
//   required: true
//   schema:
//     "$ref": "#/definitions/Student"
// responses:
//   "201":
//     "$ref": "#/responses/booleanRes"
//   "400":
//     "$ref": "#/responses/badReq"
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

// swagger:operation PUT /students Students EditStudent
// ---
// summary: Modifie un étudiant
// parameters:
// - name: student
//   description: L'étudiant à modifier
//   in: body
//   required: true
//   schema:
//     "$ref": "#/definitions/Student"
// responses:
//   "200":
//     "$ref": "#/responses/booleanRes"
//   "400":
//     "$ref": "#/responses/badReq"
//   "404":
//     "$ref": "#/responses/notFoundReq"
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

// swagger:operation DELETE /students/{id} Students DeleteStudent
// ---
// summary: Supprime l'étudiant d'id spécifié
// parameters:
// - name: id
//   in: path
//   description: id de l'étudiant
//   type: string
//   required: true
// - name: student
//   description: L'étudiant à supprimer
//   in: body
//   required: true
//   schema:
//     "$ref": "#/definitions/Student"
// responses:
//   "200":
//     "$ref": "#/responses/booleanRes"
//   "400":
//     "$ref": "#/responses/badReq"
//   "404":
//     "$ref": "#/responses/notFoundReq"
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
