package main

import (
	"fmt"
	"oops/main/infrastructure"
	"oops/main/internal"
)

var courseResults []internal.CourseResult

func main() {
	registrar := internal.Registrar{}

	registrar.LoadCourses()
	registrar.DisplayCourses()

	registrar.LoadStudents()
	registrar.DisplayStudents()

	courseResults = infrastructure.LoadCourseResults()

	fmt.Println("======================")
	fmt.Println()

	fmt.Print(courseResults)
}
