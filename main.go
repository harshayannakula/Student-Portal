package main

import (
	"oops/main/domain"
	"oops/main/export"
	"oops/main/reports"
)

func main() {
	studentRegister := domain.StudentRegister{}
	studentRegister.LoadStudents()
	studentRegister.Display()

	resultregister := domain.ResultRegister{}
	resultregister.SetResults(reports.LoadResultsForStudents(studentRegister.Students))
	resultregister.Display()

	freqMap := domain.CountGradeFrequencies(resultregister.Results())
	export.PrintHorizontalBarGraph(freqMap)
	export.ExportGPATrend(freqMap)

  
	// ------------ Boot Objects ------------
	/*
	alice := domain.NewStudent(1, "Alice Uncle")
		core := domain.NewCourse(101, "Intro to Go")
		lab := domain.NewCourse(202, "Go Networking lab")

		reg := &domain.Registrar{}
		reg.AddStudent(alice)
		reg.AddCourse(core)
		reg.AddCourse(lab)

		// Initial grading policies
		reg.Enroll(domain.NewEnrollment(alice, core, domain.PercentageGrader{}, 0.91))
		reg.Enroll(domain.NewEnrollment(alice, lab, domain.PassFailGrader{PassMark: 0.7}, 0.65))


		// Runtime policy change
		reg.SetGrader(lab.Id, domain.PercentageGrader{})

		// Prints
		for _, e := range reg.Enrollments() {
			fmt.Println(e)
		}

		// Persistence
		if err := export.ExportTranscript("transcript.csv", reg.Enrollments()); err != nil {
			fmt.Println("CSV error: ", err)
		}
	*/
}
