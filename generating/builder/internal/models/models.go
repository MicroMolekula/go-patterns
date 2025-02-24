package models

type Student struct {
	Name    string
	Surname string
	Group   string
}

type Teacher struct {
	Name    string
	Surname string
}

type FrontPageData struct {
	Student    Student
	Teacher    Teacher
	NumberWork int
}
