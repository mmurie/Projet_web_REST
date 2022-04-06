package entities

import "strconv"

type Student struct {
	Id           int
	FirstName    string
	LastName     string
	LanguageCode string
}

func NewStudent(id int, firstName string, lastName string, languageCode string) Student {
	st := new(Student)
	st.Id = id
	st.FirstName = firstName
	st.LastName = lastName
	st.LanguageCode = languageCode
	return *st
}

func (student Student) String() string {
	return "Id = " + strconv.Itoa(student.Id) + " ; " +
		"FirstName = " + student.FirstName + " ; " +
		"LastName = " + student.LastName + " ; " +
		"LanguageCode = " + student.LanguageCode
}
