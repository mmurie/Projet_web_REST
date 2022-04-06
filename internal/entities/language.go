package model

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

func String(language Language) string {
	return "Code = " + language.Code + " ; " +
		"Name = " + language.Name
}
