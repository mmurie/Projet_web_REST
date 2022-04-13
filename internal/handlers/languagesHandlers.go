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

// swagger:operation GET /languages Languages GetAllLanguages
// ---
// summary: Renvoie la liste de tous les langages de prorgammation
// responses:
//   "200":
//     "$ref": "#/responses/languagesRes"
//   "400":
//     "$ref": "#/responses/badReq"
func (lh *LanguagesHandlers) GetAllLanguages(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" - - - - - - - - - - - ")
	fmt.Println("func GetAllLanguages")

	jsonString, err := json.Marshal(lh.DAO.FindAll())
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, string(jsonString))

}

// swagger:operation GET /languages/{code} Languages GetLanguages
// ---
// summary: Renvoie le langage ayant le code spécifié
// parameters:
// - name: id
//   in: path
//   description: code du langage
//   type: string
//   required: true
// responses:
//   "200":
//     "$ref": "#/responses/languageRes"
//   "400":
//     "$ref": "#/responses/badReq"
//   "404":
//     "$ref": "#/responses/notFoundReq"
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

// swagger:operation POST /languages Languages AddLanguage
// ---
// summary: Ajoute un nouveau langage de progammation
// parameters:
// - name: language
//   description: Le langage à ajouter
//   in: body
//   required: true
//   schema:
//     "$ref": "#/definitions/Language"
// responses:
//   "201":
//     "$ref": "#/responses/booleanRes"
//   "400":
//     "$ref": "#/responses/badReq"
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

// swagger:operation PUT /languages/{code} Languages EditLanguage
// ---
// summary: Modifie un langage de progammation
// parameters:
// - name: id
//   in: path
//   description: code du langage
//   type: string
//   required: true
// - name: language
//   description: Le langage à modifier
//   in: body
//   required: true
//   schema:
//     "$ref": "#/definitions/Language"
// responses:
//   "200":
//     "$ref": "#/responses/booleanRes"
//   "400":
//     "$ref": "#/responses/badReq"
//   "404":
//     "$ref": "#/responses/notFoundReq"
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

// swagger:operation DELETE /languages/{code} Languages DeleteLanguage
// ---
// summary: Supprime un langage de progammation
// parameters:
// - name: id
//   in: path
//   description: code du langage
//   type: string
//   required: true
// - name: language
//   description: Le langage à supprimer
//   in: body
//   required: true
//   schema:
//     "$ref": "#/definitions/Language"
// responses:
//   "200":
//     "$ref": "#/responses/booleanRes"
//   "400":
//     "$ref": "#/responses/badReq"
//   "404":
//     "$ref": "#/responses/notFoundReq"
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
