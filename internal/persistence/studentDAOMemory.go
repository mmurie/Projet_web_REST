package persistence

import (
	"internal/entities"
	"sort"
)

type StudentDAOMemory struct {
	Students map[int]entities.Student
}

func (studentDAOMemory *StudentDAOMemory) FindAll() map[int]entities.Student {
	keys := make([]int, 0, len(studentDAOMemory.Students))
	for k := range studentDAOMemory.Students {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	return studentDAOMemory.Students
}

func (studentDAOMemory *StudentDAOMemory) Find(id int) entities.Student {
	return studentDAOMemory.Students[id]
}

func (studentDAOMemory *StudentDAOMemory) Exists(id int) bool {
	_, exists := studentDAOMemory.Students[id]

	return exists
}

func (studentDAOMemory *StudentDAOMemory) Delete(id int) bool {
	exists := studentDAOMemory.Exists(id)
	if exists {
		delete(studentDAOMemory.Students, id)
	}

	return exists
}

func (studentDAOMemory *StudentDAOMemory) Create(student entities.Student) bool {
	exists := studentDAOMemory.Exists(student.Id)
	if !exists {
		studentDAOMemory.Students[student.Id] = student
	}

	return !exists
}

func (studentDAOMemory *StudentDAOMemory) Update(student entities.Student) bool {
	exists := studentDAOMemory.Exists(student.Id)
	if exists {
		studentDAOMemory.Students[student.Id] = student
	}

	return exists
}
