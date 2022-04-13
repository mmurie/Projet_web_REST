package entities

// Response payload pour un étudiant
// swagger:response studentRes
type swaggStudentRes struct {
	// in:body
	Body Student
}

// Response payload pour une liste d'étudiants
// swagger:response studentsRes
type swaggStudentsRes struct {
	// in:body
	Body map[string]Student
}

// Response payload pour un langage
// swagger:response languageRes
type swaggLanguageRes struct {
	// in:body
	Body Language
}

// Response payload pour une liste de langages
// swagger:response languagesRes
type swaggLanguagesRes struct {
	// in:body
	Body map[string]Language
}

// Erreur 400 : Bad Request
// swagger:response badReq
type swaggReqBadRequest struct {
	// in:body
	Body struct {
		// HTTP status code 400 -  Bad Request
		Code int `json:"code"`
	}
}

// Erreur 404 : Not Found
// swagger:response notFoundReq
type swaggReqNotFound struct {
	// in:body
	Body struct {
		// HTTP status code 404 -  Not Found
		Code int `json:"code"`
	}
}

// True si l'opération est un succès, false sinon
// swagger:response booleanRes
type swaggBooleanRes bool
