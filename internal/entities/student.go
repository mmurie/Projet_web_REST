package entities

import "strconv"

// Un étudiant
type Student struct {
	Id           int    `json:"Id"`
	FirstName    string `json:"FirstName"`
	LastName     string `json:"LastName"`
	Age          int    `json:"Age"`
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

// Student response payload
// swagger:response studentRes
type swaggStudentRes struct {
	// in:body
	Body Student
}
