package entities

// Un langage de programmation
type Language struct {
	//Le code unique du langage
	Code string `json:"Code"`
	//Le nom du langage
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
