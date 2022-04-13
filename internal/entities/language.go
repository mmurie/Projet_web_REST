package entities

// Un language de programmation
type Language struct {
	Code string `json:"Code"`
	Name string `json:"Name"`
}

func NewLanguage(code string, name string) Language {
	lg := new(Language)
	lg.Code = code
	lg.Name = name
	return *lg
}

func (language Language) String() string {
	return "Code = " + language.Code + " ; " +
		"Name = " + language.Name
}

// Response payload pour un language
// swagger:response languageRes
type swaggLanguageRes struct {
	// in:body
	Body Language
}

// Response payload pour une liste de languages
// swagger:response languagesRes
type swaggLanguagesRes struct {
	// in:body
	Body []Language
}
