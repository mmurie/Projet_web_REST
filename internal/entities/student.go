package entities

import "strconv"

type Student struct {
	Id           int
	FirstName    string
	LastName     string
	Age          int
	LanguageCode string
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
