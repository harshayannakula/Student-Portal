package admin

import (
	"encoding/json"
	"fmt"
	"log"
	"oops/main/internal"

	//"oops/main/analytics"
	//"oops/main/domain"
	"os"
)

type Registrar struct {
	students    []internal.Student
	courses     []internal.Course
	enrollments []internal.Enrollment
	graders     map[int]internal.Grader
}

func NewRegistrar() *Registrar {
	return &Registrar{graders: make(map[int]internal.Grader)}
}

func (r *Registrar) AddStudent(s internal.Student) {
	r.students = append(r.students, s)
}

func (r *Registrar) AddCourse(c internal.Course) {
	r.courses = append(r.courses, c)
}

func (r *Registrar) Enroll(e internal.Enrollment) {
	r.enrollments = append(r.enrollments, e)
}

func (r *Registrar) SetGrader(courseID int, g internal.Grader) {
	for i, e := range r.enrollments {
		if e.Course.Id == courseID {
			r.enrollments[i].Grader = g
		}
	}
}

func (r *Registrar) Enrollments() []internal.Enrollment {
	return r.enrollments
}

type StudentData struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type courseData struct {
	ID      int     `json:"id"`
	Name    string  `json:"title"`
	Credits float64 `json:"credits"`
}

func (regis *Registrar) LoadCourses() {
	var coursesData []courseData
	data, err := os.ReadFile("courses.json")
	if err != nil {
		log.Fatal("Failed to read courses.json:", err)
	}
	err = json.Unmarshal(data, &coursesData)
	if err != nil {
		log.Fatal("Failed to unmarshal courses:", err)
	}
	for _, course := range coursesData {
		regis.AddCourse(internal.NewCourse(course.ID, course.Name))
	}
}

func (r *Registrar) LoadStudents() {
	var students []StudentData
	data, err := os.ReadFile("students.json")
	if err != nil {
		log.Fatal("Failed to read students.json:", err)
	}

	err = json.Unmarshal(data, &students)
	if err != nil {
		log.Fatal("Failed to unmarshal students:", err)
	}
	r.students = make([]internal.Student, 0, len(students))
	for _, sd := range students {
		student := internal.NewStudent(sd.ID, sd.Name)
		r.AddStudent(student)
	}
}

func (r *Registrar) DisplayStudents() {
	for _, sd := range r.students {
		fmt.Printf("#%d : %s\n", sd.ID(), sd.Name())
	}
}

func (r *Registrar) DisplayCourses() {
	log.Println("Courses...")
	for _, cr := range r.courses {
		log.Printf("#%d : %s\n", cr.Id, cr.Name)
	}
}
