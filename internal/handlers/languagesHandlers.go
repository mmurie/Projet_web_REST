package handlers

import (
	"encoding/json"
	"fmt"
	"internal/persistence"
	"net/http"

	"github.com/gorilla/mux"
)

type LanguagesHandlers struct {
	DAO persistence.LanguageDAO
}

func (lh *LanguagesHandlers) GetAllLanguages(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" - - - - - - - - - - - ")
	fmt.Println("func GetAllLanguages")

	jsonString, err := json.Marshal(lh.DAO.FindAll())
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, string(jsonString))

}

func (lh *LanguagesHandlers) GetLanguage(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" - - - - - - - - - - - ")
	fmt.Println("func GetLanguage")

	vars := mux.Vars(r)
	code := vars["code"]

	jsonString, err := json.Marshal(lh.DAO.Find(code))
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

func (lh *LanguagesHandlers) InitializeLanguagesRoutes(r *mux.Router) {
	r.HandleFunc("/rest/languages/{code}", lh.GetLanguage).Methods("GET")
	r.HandleFunc("/rest/languages", lh.GetAllLanguages).Methods("GET")
	r.HandleFunc("/rest/languages", AddLanguage).Methods("POST")
	r.HandleFunc("/rest/languages", EditLanguage).Methods("PUT")
	r.HandleFunc("/rest/languages/{code}", DeleteLanguage).Methods("DELETE")
}
