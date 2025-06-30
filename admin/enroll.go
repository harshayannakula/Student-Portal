package admin

import (
	"fmt"
)

type Enrollment struct {
	Student
	Course
	Grader
	score float64
}

type Enrollnew struct {
	Enrollment              // Embedding Enrollment to include student, course, grader, and score
	attend     []Attendance // Attendance records associated with the enrollment
	Teacher                 // Embedding Teacher to associate with the enrollment
}

type Teacherenrollment struct {
	Teacher
	Course
}

func NewTeacherenrollment(t Teacher, c Course) Teacherenrollment {
	return Teacherenrollment{Teacher: t, Course: c}
}

func NewEnrollment(st Student, c Course, g Grader, score float64) Enrollment {
	return Enrollment{Student: st, Course: c, Grader: g, score: score}
}

func enrollnew(st Student, c Course, g Grader, score float64, att []Attendance, t Teacher) Enrollnew {
	return Enrollnew{Enrollment: Enrollment{Student: st, Course: c, Grader: g, score: score}, attend: att}
}

func (e Enrollment) String() string {
	r, _ := e.Grade(e)
	return fmt.Sprintf("%s ->  %s : %s", e.Student.Name(), e.Course.Name, r)
}
