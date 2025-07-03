package main

import (
	"fmt"
	"oops/main/infrastructure"
	"oops/main/internal"
)

var courseResults []internal.CourseResult

func main() {
	regitrar := internal.Registrar{}

	regitrar.LoadCourses()
	regitrar.DisplayCourses()

	regitrar.LoadStudents()
	regitrar.DisplayStudents()
	
	courseResults = infrastructure.LoadCourseResults()

	fmt.Println("======================")
	fmt.Println()

	fmt.Print(courseResults)
}

