package internal

type Course struct {
	Id   int
	Name string
}

type CreditCourse struct {
	Course
	Credit int
}

func NewCreditCourse(id int, name string, credits int) CreditCourse {
	return CreditCourse{Course: Course{Id: id, Name: name}, Credit: credits}

}

func NewCourse(id int, name string) Course {
	return Course{Id: id, Name: name}
}
