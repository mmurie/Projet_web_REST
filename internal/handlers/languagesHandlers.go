package handlers

import (
	"encoding/json"
	"fmt"
	"internal/hardcodedData"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllLanguages(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" - - - - - - - - - - - ")
	fmt.Println("func getAllLanguages")

	jsonString, err := json.Marshal(hardcodedData.GetAllDatas())
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, string(jsonString))

}

func GetLanguage(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" - - - - - - - - - - - ")
	fmt.Println("func getLanguage")

	vars := mux.Vars(r)
	code := vars["code"]

	jsonString, err := json.Marshal(hardcodedData.GetData(code))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, string(jsonString))
}

func AddLanguage(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("func addLanguage")
}

func EditLanguage(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("func editLanguage")
}

func DeleteLanguage(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("func getLanguage")
}

func InitializeLanguageRoutes(r *mux.Router) {
	r.HandleFunc("/rest/languages/{code}", GetLanguage).Methods("GET")
	r.HandleFunc("/rest/languages", GetAllLanguages).Methods("GET")
	r.HandleFunc("/rest/languages", AddLanguage).Methods("POST")
	r.HandleFunc("/rest/languages", EditLanguage).Methods("PUT")
	r.HandleFunc("/rest/languages/{code}", DeleteLanguage).Methods("DELETE")

}
