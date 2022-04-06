package persistence

import (
	"internal/entities"
	"sort"
)

type LanguageDAOMemory struct {
	Languages map[string]entities.Language
}

func (languageDAOMemory *LanguageDAOMemory) FindAll() map[string]entities.Language {
	keys := make([]string, 0, len(languageDAOMemory.Languages))
	for k := range languageDAOMemory.Languages {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return languageDAOMemory.Languages
}

func (languageDAOMemory *LanguageDAOMemory) Find(code string) entities.Language {
	return languageDAOMemory.Languages[code]
}

func (languageDAOMemory *LanguageDAOMemory) Exists(code string) bool {
	_, exists := languageDAOMemory.Languages[code]

	return exists
}

func (languageDAOMemory *LanguageDAOMemory) Delete(code string) bool {
	exists := languageDAOMemory.Exists(code)
	if exists {
		delete(languageDAOMemory.Languages, code)
	}

	return exists
}

func (languageDAOMemory *LanguageDAOMemory) Create(language entities.Language) bool {
	exists := languageDAOMemory.Exists(language.Code)
	if !exists {
		languageDAOMemory.Languages[language.Code] = language
	}

	return !exists
}

func (languageDAOMemory *LanguageDAOMemory) Update(language entities.Language) bool {
	exists := languageDAOMemory.Exists(language.Code)
	if exists {
		languageDAOMemory.Languages[language.Code] = language
	}

	return exists
}
