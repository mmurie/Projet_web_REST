package handlers

import (
	"encoding/json"
	"fmt"
	"internal/persistence"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllLanguages(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" - - - - - - - - - - - ")
	fmt.Println("func GetAllLanguages")

	jsonString, err := json.Marshal( /*languageDAO.FindAll()*/ "")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, string(jsonString))

}

func GetLanguage(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" - - - - - - - - - - - ")
	fmt.Println("func GetLanguage")

	vars := mux.Vars(r)
	code := vars["code"]

	jsonString, err := json.Marshal( /*languageDAO.Find(code)*/ code)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, string(jsonString))
}

func AddLanguage(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" - - - - - - - - - - - ")
	fmt.Println("func AddLanguage")
}

func EditLanguage(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" - - - - - - - - - - - ")
	fmt.Println("func EditLanguage")
}

func DeleteLanguage(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" - - - - - - - - - - - ")
	fmt.Println("func GetLanguage")
}

func InitializeLanguagesRoutes(r *mux.Router, languageDAO *persistence.LanguageDAO) {
	r.HandleFunc("/rest/languages/{code}", GetLanguage).Methods("GET")
	r.HandleFunc("/rest/languages", GetAllLanguages).Methods("GET")
	r.HandleFunc("/rest/languages", AddLanguage).Methods("POST")
	r.HandleFunc("/rest/languages", EditLanguage).Methods("PUT")
	r.HandleFunc("/rest/languages/{code}", DeleteLanguage).Methods("DELETE")

}
