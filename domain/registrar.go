package domain

type Registrar struct {
	students    []Student
	courses     []Course
	enrollments []Enrollment
}

func (r *Registrar) AddStudent(s Student) {
	r.students = append(r.students, s)
}

func (r *Registrar) AddCourse(c Course) {
	r.courses = append(r.courses, c)
}

func (r *Registrar) Enroll(e Enrollment) {
	r.enrollments = append(r.enrollments, e)
}


func (r *Registrar) SetGrader(courseID int, g Grader){
	for i, e := range r.enrollments {
		if e.Course.Id == courseID {
			r.enrollments[i].Grader = g
		}
	}
}

func (r *Registrar) Enrollments() []Enrollment {
	return r.enrollments
}
