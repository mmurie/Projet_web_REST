package persistence

import (
	"encoding/json"
	"internal/entities"
	"log"
	"sort"
)

type LanguageDAOBolt struct {
	Dbms BoltDb
}

func (languageDAOBolt *LanguageDAOBolt) FindAll() map[string]entities.Language {
	languagesStrings := languageDAOBolt.Dbms.DbGetAll("languages")
	languages := make(map[string]entities.Language)

	for _, element := range languagesStrings {
		language := new(entities.Language)
		decode_err := json.Unmarshal([]byte(element), &language)
		if decode_err != nil {
			log.Fatal(decode_err)
		}
		languages[language.Code] = *language
	}

	keys := make([]string, 0, len(languages))
	for k := range languages {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return languages
}

func (languageDAOBolt *LanguageDAOBolt) Find(code string) entities.Language {
	languageString := languageDAOBolt.Dbms.DbGet("languages", code)
	language := new(entities.Language)
	decode_err := json.Unmarshal([]byte(languageString), &language)
	if decode_err != nil {
		log.Fatal(decode_err)
	}
	return *language
}

func (languageDAOBolt *LanguageDAOBolt) Exists(code string) bool {
	exists := languageDAOBolt.Find(code) != entities.Language{}

	return exists
}

func (languageDAOBolt *LanguageDAOBolt) Delete(code string) bool {
	exists := languageDAOBolt.Exists(code)
	if exists {
		languageDAOBolt.Dbms.DbDelete("languages", code)
	}

	return exists
}

func (languageDAOBolt *LanguageDAOBolt) Create(language entities.Language) bool {
	exists := languageDAOBolt.Exists(language.Code)
	if !exists {
		encoded, encoded_err := json.Marshal(language)
		if encoded_err != nil {
			log.Fatal(encoded_err)
		}
		languageDAOBolt.Dbms.DbPut("languages", language.Code, string(encoded))
	}

	return !exists
}

func (languageDAOBolt *LanguageDAOBolt) Update(language entities.Language) bool {
	exists := languageDAOBolt.Exists(language.Code)
	if exists {
		encoded, encoded_err := json.Marshal(language)
		if encoded_err != nil {
			log.Fatal(encoded_err)
		}
		languageDAOBolt.Dbms.DbPut("languages", language.Code, string(encoded))
	}

	return exists
}
