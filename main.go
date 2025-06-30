package main

import (
	"fmt"
	"oops/main/domain"
	"oops/main/reports"
)

var courseResults []domain.CourseResult

func main() {
	regitrar := domain.Registrar{}

	regitrar.LoadCourses()
	regitrar.DisplayCourses()

	regitrar.LoadStudents()
	regitrar.DisplayStudents()

	courseResults = reports.LoadCourseResults()

	fmt.Println("======================")
	fmt.Println()

	fmt.Print(courseResults)
}
