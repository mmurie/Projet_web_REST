package entities

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
