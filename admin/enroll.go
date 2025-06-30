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

type Enroll struct {
	Enrollment
	attend []Attendance
}

func Updateattendance(e Enroll, a Attendance) Enroll {
	e.attend = append(e.attend, a)
	return e
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

func (e Enrollment) String() string {
	r, _ := e.Grade(e)
	return fmt.Sprintf("%s ->  %s : %s", e.Student.Name(), e.Course.Name, r)
}
