package internal

import (
	"os"
	"testing"
	"time"
)

func createMockRegistrar() NewRegistrar {
	student := NewStudent(2, "Bob")
	course := NewCourse(202, "Python")
	creditCourse := NewCreditCourse(course, 4)
	teacher := Teacher{id: 2, Name: "Prof. Smith"}
	teacherenrollment := NewTeacherenrollment(teacher, creditCourse)
	enrollment := NewEnrollment(student, course, nil, 0.0)
	attendance := Attendance{
		Records: map[time.Time]bool{
			time.Date(2023, 11, 1, 0, 0, 0, 0, time.UTC): true,
		},
	}
	enrollNew, _ := Enroll(enrollment, attendance, teacher, []Teacherenrollment{teacherenrollment})

	r := NewRegistrar{}
	r.AddStudent(student)
	r.AddCourse(course)
	r.AddTeacher(teacher)
	r.AddTeacherenrollment(teacherenrollment)
	r.Enroll(enrollment)
	r.AddEnrollnew(enrollNew)

	return r
}

func TestAddFunctions(t *testing.T) {
	reg := NewRegistrar{}

	s := NewStudent(3, "John")
	c := NewCourse(301, "Java")
	cc := NewCreditCourse(c, 3)
	te := Teacher{id: 3, Name: "TeacherX"}
	tenroll := NewTeacherenrollment(te, cc)
	enroll := NewEnrollment(s, c, nil, 85.0)
	att := Attendance{Records: map[time.Time]bool{time.Now(): true}}

	reg.AddStudent(s)
	reg.AddCourse(c)
	reg.AddTeacher(te)
	reg.AddTeacherenrollment(tenroll)
	reg.Enroll(enroll)

	enew, ok := Enroll(enroll, att, te, []Teacherenrollment{tenroll})
	if !ok {
		t.Error("Expected successful enrollment with teacher mapping")
	}
	reg.AddEnrollnew(enew)
}

func TestEnrollFailure(t *testing.T) {
	student := NewStudent(4, "Fail Student")
	course := NewCourse(404, "FailCourse")
	enroll := NewEnrollment(student, course, nil, 0.0)
	att := Attendance{}
	teacher := Teacher{id: 99, Name: "Unassigned"}

	enew, ok := Enroll(enroll, att, teacher, []Teacherenrollment{}) // No match
	if ok {
		t.Error("Expected failure due to unmatched teacher mapping")
	}
	if enew.Teacher.ID() != 0 {
		t.Error("Expected empty Enrollnew struct")
	}
}

func TestGiveAndFetchAttendance(t *testing.T) {
	reg := createMockRegistrar()
	now := time.Now()

	// Valid attendance marking
	ok := Giveattendence(&reg, 202, 2, 2, true, now)
	if !ok {
		t.Error("Giveattendence failed on valid data")
	}

	records, ok := FetchAttendance(&reg, 202, 2, 2)
	if !ok {
		t.Error("FetchAttendance failed on valid data")
	}
	if val, exists := records[now]; !exists || !val {
		t.Error("Attendance value incorrect or not set")
	}
}

func TestGiveAttendanceInvalid(t *testing.T) {
	reg := createMockRegistrar()
	ok := Giveattendence(&reg, 999, 999, 999, true, time.Now())
	if ok {
		t.Error("Expected Giveattendence to fail on invalid IDs")
	}
}

func TestFetchAttendanceInvalid(t *testing.T) {
	reg := createMockRegistrar()
	_, ok := FetchAttendance(&reg, 999, 999, 999)
	if ok {
		t.Error("Expected FetchAttendance to fail on invalid data")
	}
}

func TestStudentPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for invalid student ID")
		}
	}()
	_ = NewStudent(0, "Invalid") // Should panic
}

func TestDisplayFunctions(t *testing.T) {
	reg := NewRegistrar{}
	reg.AddStudent(NewStudent(6, "DisplayTest"))
	reg.AddCourse(NewCourse(606, "DisplayCourse"))

	// Just for coverage; visually nothing is tested
	reg.DisplayStudents()
	reg.DisplayCourses()
}

func TestLoadStudentsAndCourses(t *testing.T) {
	// Write temporary JSON files
	studentJSON := `[{"id":7,"name":"JSON Student"}]`
	courseJSON := `[{"id":707,"title":"JSON Course","credits":3}]`

	_ = os.WriteFile("students.json", []byte(studentJSON), 0644)
	_ = os.WriteFile("courses.json", []byte(courseJSON), 0644)

	reg := &Registrar{}
	reg.LoadStudents()
	reg.LoadCourses()

	if len(reg.students) == 0 || len(reg.courses) == 0 {
		t.Error("Failed to load students or courses from JSON")
	}

	// Clean up
	_ = os.Remove("students.json")
	_ = os.Remove("courses.json")
}
