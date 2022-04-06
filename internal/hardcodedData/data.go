package hardcodedData

import (
	"internal/entities"
)

func GetAllDatas() map[string]entities.Language {
	language1 := entities.NewLanguage("js", "JavaScript")
	language2 := entities.NewLanguage("c", "C")

	languages := make(map[string]entities.Language)
	languages["js"] = language1
	languages["c"] = language2

	return languages
}

func GetData(code string) entities.Language {
	language1 := entities.NewLanguage("js", "JavaScript")
	language2 := entities.NewLanguage("c", "C")

	languages := make(map[string]entities.Language)
	languages["js"] = language1
	languages["c"] = language2

	return languages[code]
}
