package internal

// import (
// 	"testing"
// 	"time"
// )

// type Teacher struct {
// 	Name string
// }

// func Testenrollnew(t *testing.T) {
// 	// Parse the date string to time.Time
// 	date, err := time.Parse("2006-01-02", "2023-10-01")
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Correct initialization
// 	attend := Attendance{
// 		Records: map[time.Time]bool{
// 			date: true,
// 		},
// 	}
// 	student := NewStudent(1, "Alice")
// 	course := NewCourse(101, "GoLang")
// 	grade := GradeLetter{}
// 	teacher := Teacher{ID: "01", Name: "Mr. Smith"}
// 	enroll := Enrollnew(student, course, grade, 90, attend, teacher)

// 	if enroll.Student.Name() != "Alice" || enroll.Course.Name != "GoLang" {
// 		t.Errorf("Enrollment creation failed")
// 	}
// }

// func TestTeacherenrollment(t *testing.T) {
// 	teacher := Teacher{ID: "01", Name: "Mr. Smith"}
// 	//course := NewCourse(101, "GoLang")
// 	ccourse := NewCreditCourse(102, "Python", 3)

// 	te := NewTeacherenrollment(teacher, ccourse)

// 	if te.Teacher.Name() != "Mr. John" {
// 		t.Errorf("Teacher enrollment failed")
// 	}
// }
