package internal

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
	Enrollment            // Embedding Enrollment to include student, course, grader, and score
	attend     Attendance // Attendance records associated with the enrollment
	Teacher               // Embedding Teacher to associate with the enrollment
}

// teacher with course
type Teacherenrollment struct {
	Teacher
	CreditCourse
}

// function to map teacher with course
func NewTeacherenrollment(t Teacher, c CreditCourse) Teacherenrollment {
	return Teacherenrollment{Teacher: t, CreditCourse: c}
}

// old function without teacher ,attandence
func NewEnrollment(st Student, c Course, g Grader, score float64) Enrollment {
	return Enrollment{Student: st, Course: c, Grader: g, score: score}
}

// made enrool function to check with teacher maping with course (changes may be like checking in main while adding or passing teacher map from rigister to enroll function)
// func Enroll(st Student, c Course, g Grader, score float64, att Attendance, t Teacher, teachermap []Teacherenrollment) (Enrollnew, bool) {
func Enroll(oe Enrollment, att Attendance, t Teacher, teachermap []Teacherenrollment) (Enrollnew, bool) {
	for _, te := range teachermap {
		if te.Teacher.ID() == t.ID() && te.Course.Id == oe.Course.Id {
			return Enrollnew{Enrollment: oe, attend: att, Teacher: t}, true
		}
	}
	return Enrollnew{}, false
}

func (e Enrollment) String() string {
	r, _ := e.Grade(e)
	return fmt.Sprintf("%s ->  %s : %s", e.Student.Name(), e.Course.Name, r)
}
