package domain

type Course struct {
	Id   int
	Name string
}

func NewCourse(id int, name string) Course {
	return Course{Id: id, Name: name}
}
