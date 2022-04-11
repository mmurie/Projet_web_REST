package persistence

import (
	"encoding/json"
	"internal/entities"
	"log"
	"sort"
	"strconv"
)

type StudentDAOBolt struct {
	Dbms BoltDb
}

func (studentDAOBolt *StudentDAOBolt) FindAll() map[int]entities.Student {
	studentsStrings := studentDAOBolt.Dbms.DbGetAll("students")
	students := make(map[int]entities.Student)

	for _, element := range studentsStrings {
		student := new(entities.Student)
		decode_err := json.Unmarshal([]byte(element), &student)
		if decode_err != nil {
			log.Fatal(decode_err)
		}
		students[student.Id] = *student
	}

	keys := make([]int, 0, len(students))
	for k := range students {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	return students
}

func (studentDAOBolt *StudentDAOBolt) Find(id int) entities.Student {
	studentString := studentDAOBolt.Dbms.DbGet("students", strconv.Itoa(id))
	student := new(entities.Student)
	decode_err := json.Unmarshal([]byte(studentString), &student)
	if decode_err != nil {
		//log.Fatal(decode_err)
	}
	return *student
}

func (studentDAOBolt *StudentDAOBolt) Exists(id int) bool {
	exists := studentDAOBolt.Find(id) != entities.Student{}

	return exists
}

func (studentDAOBolt *StudentDAOBolt) Delete(id int) bool {
	exists := studentDAOBolt.Exists(id)
	if exists {
		studentDAOBolt.Dbms.DbDelete("students", strconv.Itoa(id))
	}

	return exists
}

func (studentDAOBolt *StudentDAOBolt) Create(student entities.Student) bool {
	exists := studentDAOBolt.Exists(student.Id)
	if !exists {
		encoded, encoded_err := json.Marshal(student)
		if encoded_err != nil {
			log.Fatal(encoded_err)
		}
		studentDAOBolt.Dbms.DbPut("students", strconv.Itoa(student.Id), string(encoded))
	}

	return !exists
}

func (studentDAOBolt *StudentDAOBolt) Update(student entities.Student) bool {
	exists := studentDAOBolt.Exists(student.Id)
	if exists {
		encoded, encoded_err := json.Marshal(student)
		if encoded_err != nil {
			log.Fatal(encoded_err)
		}
		studentDAOBolt.Dbms.DbPut("students", strconv.Itoa(student.Id), string(encoded))
	}

	return exists
}
