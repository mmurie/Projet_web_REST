package entities

import "strconv"

// Un étudiant
type Student struct {
	//L'identifiant unique de l'étudiant
	Id int `json:"Id"`
	//Le prénom de l'étudiant
	FirstName string `json:"FirstName"`
	//Le nom de l'étudiant
	LastName string `json:"LastName"`
	//L'âge de l'étudiant
	Age int `json:"Age"`
	//Le code du langage de programmation de l'étudiant
	LanguageCode string `json:"LanguageCode"`
}

func NewStudent(id int, firstName string, lastName string, age int, languageCode string) Student {
	st := new(Student)
	st.Id = id
	st.FirstName = firstName
	st.LastName = lastName
	st.Age = age
	st.LanguageCode = languageCode
	return *st
}

func (student Student) String() string {
	return "Id = " + strconv.Itoa(student.Id) + " ; " +
		"FirstName = " + student.FirstName + " ; " +
		"LastName = " + student.LastName + " ; " +
		"Age = " + strconv.Itoa(student.Age) + " ; " +
		"LanguageCode = " + student.LanguageCode
}

// Response payload pour un étudiant
// swagger:response studentRes
type swaggStudentRes struct {
	// in:body
	Body Student
}

// Response payload pour une liste d'étudiants
// swagger:response studentsRes
type swaggStudentsRes struct {
	// in:body
	Body []Student
}
