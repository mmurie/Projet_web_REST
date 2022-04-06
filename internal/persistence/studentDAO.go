package persistence

import (
	"internal/entities"
)

type StudentDAO interface {
	FindAll() map[string]entities.Student //triés par Id
	Find(id int) entities.Student
	Exists(id int) bool
	Delete(id int) bool                   //false si non trouvé
	Create(student entities.Student) bool //false si existe déjà
	Update(student entities.Student) bool //false si non trouvé
}
