package internal

import "testing"

var student1 = NewStudent(1, "alice")

var course1 = NewCourse(101, "Golang")

//var Registrar1 = NewRegistrar()

func TestAddStudent(t *testing.T) {
	Registrar1 := NewRegistrar{
		Registrar: Registrar{
			students:    []Student{},
			courses:     []Course{},
			enrollments: []Enrollment{},
		},
		teacher:    []Teacher{},
		teachermap: []Teacherenrollment{},
		enroll:     []Enrollnew{},
	}

	Registrar1.AddStudent(student1)
	if len(Registrar1.students) != 1 {
		t.Errorf("Expected 1 student, got %d", len(Registrar1.students))
	}
	if Registrar1.students[0].Name() != "alice" {
		t.Errorf("Expected student name 'alice', got '%s'", Registrar1.students[0].Name())
	}
}
