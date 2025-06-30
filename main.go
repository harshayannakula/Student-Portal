package main

import (
	"fmt"
	"oops/main/internal/students"
	"oops/main/infrastructure/reports"
	"oops/main/internal/admin"
)

var courseResults []students.CourseResult

func main() {
	regitrar := admin.Registrar{}

	regitrar.LoadCourses()
	regitrar.DisplayCourses()

	regitrar.LoadStudents()
	regitrar.DisplayStudents()
	
	courseResults = reports.LoadCourseResults()

	fmt.Println("======================")
	fmt.Println()

	fmt.Print(courseResults)
}

