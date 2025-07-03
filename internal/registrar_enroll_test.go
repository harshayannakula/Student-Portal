package internal

import (
	"testing"
	"time"
)

var student1 = NewStudent(1, "alice")

var course1 = NewCourse(101, "Golang")

var creditcourse1 = NewCreditCourse(course1, 3)

var teacher1 = Teacher{
	id:   1,
	Name: "bob",
}
var teacherenrollment1 = NewTeacherenrollment(teacher1, creditcourse1)

var Attendance1 = Attendance{
	Records: map[time.Time]bool{
		time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC): true,
		time.Date(2023, 10, 2, 0, 0, 0, 0, time.UTC): false,
	},
}

var Enrollment1 = NewEnrollment(student1, course1, nil, 0.0)

var Registrar1 = NewRegistrar{
	Registrar: Registrar{
		students:    []Student{},
		courses:     []Course{},
		enrollments: []Enrollment{},
	},
	teacher:    []Teacher{},
	Teachermap: []Teacherenrollment{},
	enroll:     []Enrollnew{},
}

func TestNewRegistrar(t *testing.T) {
	Registrar1.AddStudent(student1)
	Registrar1.AddCourse(course1)
	Registrar1.AddTeacher(teacher1)
	Registrar1.AddTeacherenrollment(teacherenrollment1)
	Registrar1.Enroll(Enrollment1)
	Enrollnew1, _ := Enroll(Enrollment1, Attendance1, teacher1, Registrar1.Teachermap)
	Registrar1.AddEnrollnew(Enrollnew1)

	if Registrar1.students == nil {
		t.Error("Expected students slice to be initialized, got nil")
	}
	if Registrar1.courses == nil {
		t.Error("Expected courses slice to be initialized, got nil")
	}
	if Registrar1.enrollments == nil {
		t.Error("Expected enrollments slice to be initialized, got nil")
	}
	if Registrar1.teacher == nil {
		t.Error("Expected teacher slice to be initialized, got nil")
	}
	if Registrar1.Teachermap == nil {
		t.Error("Expected teachermap slice to be initialized, got nil")
	}
	if Registrar1.enroll == nil {
		t.Error("Expected enroll slice to be initialized, got nil")
	}
}

/*
func TestAddStudent(t *testing.T) {
	Registrar1.AddStudent(student1)
	if len(Registrar1.students) != 1 {
		t.Errorf("Expected 1 student, got %d", len(Registrar1.students))
	}
	if Registrar1.students[0].Name() != "alice" {
		t.Errorf("Expected student name 'alice', got '%s'", Registrar1.students[0].Name())
	}
}

func TestAddCourse(t *testing.T) {
	Registrar1.AddCourse(course1)
	if len(Registrar1.courses) != 1 {
		t.Errorf("Expected 1 course, got %d", len(Registrar1.courses))
	}
	if Registrar1.courses[0].Name() != "Golang" {
		t.Errorf("Expected course name 'Golang', got '%s'", Registrar1.courses[0].Name())
	}
}

func TestAddTeacherenrollment(t *testing.T){
	Register1.AddTeacherenrollment(teacherenrollment1)
	if len(Registrar1.teachermap) != 1 {
		t.Errorf("Expected 1 teacher enrollment, got %d", len(Registrar1.teachermap))
	}
	if Registrar1.teachermap[0].Teacher.Name != "bob" {
}

}*/
