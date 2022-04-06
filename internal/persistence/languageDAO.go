package persistence

import (
	"internal/entities"
)

type LanguageDAO interface {
	FindAll() map[string]entities.Language //triés par Code
	Find(code string) entities.Language
	Exists(code string) bool
	Delete(code string) bool                //false si non trouvé
	Create(language entities.Language) bool //false si existe déjà
	Update(language entities.Language) bool //false si non trouvé
}
