package handlers

import (
	"encoding/json"
	"fmt"
	"internal/entities"
	"internal/persistence"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type LanguagesHandlers struct {
	DAO persistence.LanguageDAO
}

// swagger:operation GET /languages languages GetAllLanguages
// ---
// summary: Renvoie la liste de tous les languages de prorgammation
// responses:
//   "200":
//     "$ref": "#/responses/languageRes"
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

func (lh *LanguagesHandlers) AddLanguage(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" - - - - - - - - - - - ")
	fmt.Println("func AddLanguage")

	language := entities.NewLanguage("", "")
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &language)

	isOk, err := json.Marshal(lh.DAO.Create(language))

	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, string(isOk))
}

func (lh *LanguagesHandlers) EditLanguage(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" - - - - - - - - - - - ")
	fmt.Println("func EditLanguage")

	language := entities.NewLanguage("", "")
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &language)

	isOk, err := json.Marshal(lh.DAO.Update(language))

	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, string(isOk))

}

func (lh *LanguagesHandlers) DeleteLanguage(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" - - - - - - - - - - - ")
	fmt.Println("func DeleteLanguage")

	vars := mux.Vars(r)
	code := vars["code"]

	isOk, err := json.Marshal(lh.DAO.Delete(code))

	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, string(isOk))

}

func (lh *LanguagesHandlers) InitializeLanguagesRoutes(r *mux.Router) {
	r.HandleFunc("/rest/languages/{code}", lh.GetLanguage).Methods("GET")
	r.HandleFunc("/rest/languages", lh.GetAllLanguages).Methods("GET")
	r.HandleFunc("/rest/languages", lh.AddLanguage).Methods("POST")
	r.HandleFunc("/rest/languages", lh.EditLanguage).Methods("PUT")
	r.HandleFunc("/rest/languages/{code}", lh.DeleteLanguage).Methods("DELETE")
}
