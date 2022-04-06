package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllLanguages(w http.ResponseWriter, r *http.Request) {
	fmt.Println("func getAllLanguages")
}

func GetLanguage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("func getLanguage")
}

func AddLanguage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("func addLanguage")
}

func EditLanguage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("func editLanguage")
}

func DeleteLanguage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("func getLanguage")
}

func InitializeLanguageRoutes(r *mux.Router) {
	r.HandleFunc("/rest/languages/{code}", GetLanguage).Methods("GET")
	r.HandleFunc("/rest/languages", GetAllLanguages).Methods("GET")
	r.HandleFunc("/rest/languages", AddLanguage).Methods("POST")
	r.HandleFunc("/rest/languages", EditLanguage).Methods("PUT")
	r.HandleFunc("/rest/languages/{code}", DeleteLanguage).Methods("DELETE")

}
