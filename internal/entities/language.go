package entities

type Language struct {
	Code string
	Name string
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
