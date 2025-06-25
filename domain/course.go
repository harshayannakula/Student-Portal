package domain

type Course struct {
	ID   int
	name string
}

func NewCourse(id int, title string) Course {
	return Course{ID: id, name: title}
}
