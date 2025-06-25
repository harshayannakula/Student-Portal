package main

import (
	"fmt"

	"example.com/m/v3/export"
	"example.com/m/v3/reports"

	"example.com/m/v3/domain"
)

func main() {
	// st := NewStudent(1, "Alice Johson")
	// c := NewCourse(101, "Intro to Go")
	// st.Display()

	// en1 := NewEnrollment(st, c, PercentageGrader{}, 0.83)
	// en2 := NewEnrollment(st, c, PassFailGrader{passMark: 0.5}, 0.83)

	// fmt.Println(en1)
	// fmt.Println(en2)

	alice := domain.NewStudent(1, "Alice Johnson")
	core := domain.NewCourse(101, "Intro to Go")
	lab := domain.NewCourse(202, "Go Networking Lab")

	reg := &domain.Registrar{}
	reg.AddStudent(alice)
	reg.AddCourse(core)
	reg.AddCourse(lab)

	reg.Enroll(domain.NewEnrollment(alice, core, domain.PercentageGrader{}, 0.91))
	reg.Enroll(domain.NewEnrollment(alice, lab, domain.PassFailGrader{PassMark: 0.7}, 0.65))

	reg.SetGrader(lab.ID, domain.PercentageGrader{})

	for _, e := range reg.Enrollments() {
		fmt.Println(e)
	}
	if err := domain.ExportTranscript("transcript.csv", reg.Enrollments()); err != nil {
		fmt.Println("csv error.", err)
	}

	fmt.Println("Student Grade System")
	fmt.Println("======================")

	reportGen := reports.NewGPAReportGenerator()
	fmt.Println(" Generating student data...")
	reportGen.GenerateStudentData()

	students := reportGen.GetStudents()
	atRiskStudents := reportGen.GetAtRiskStudents()
	deanListStudents := reportGen.GetDeanListStudents()
	stats := reportGen.GenerateStatistics()

	fmt.Printf("\n Summary Statistics:\n")
	fmt.Printf("Total Students: %v\n", stats["total_students"])
	fmt.Printf("At Risk Students: %v (%.1f%%)\n", stats["at_risk_count"], stats["at_risk_percentage"])
	fmt.Printf("Dean's List Students: %v (%.1f%%)\n", stats["dean_list_count"], stats["dean_list_percentage"])
	fmt.Printf("Normal Students: %v\n", stats["normal_count"])
	fmt.Printf("Average GPA: %.2f/10\n", stats["average_gpa"])
	fmt.Printf("At Risk Students: %d\n", len(atRiskStudents))
	fmt.Printf("Dean's List Students: %d\n", len(deanListStudents))

	fmt.Println("\n Generating reports...")

	reports := []struct {
		filename    string
		function    func() error
		description string
	}{
		{"detailed_transcript.csv", func() error { return export.ExportDetailedTranscript("detailed_transcript.csv", students) }, "Detailed transcript"},
		{"summary_report.csv", func() error { return export.ExportSummaryReport("summary_report.csv", students) }, "Summary report"},
		{"at_risk_students.csv", func() error { return export.ExportAtRiskStudents("at_risk_students.csv", students) }, "At-risk students"},
		{"dean_list_students.csv", func() error { return export.ExportDeanListStudents("dean_list_students.csv", students) }, "Dean's list students"},
		{"statistics.csv", func() error { return export.ExportStatistics("statistics.csv", stats) }, "Statistics summary"},
	}

	for _, report := range reports {
		if err := report.function(); err != nil {
			fmt.Printf("Error generating %s: %v\n", report.description, err)
		} else {
			fmt.Printf(" %s exported to %s\n", report.description, report.filename)
		}
	}

	fmt.Println("\n All reports generated successfully!")

	fmt.Println("\n Sample Students:")
	for i, student := range students[:5] { // Show first 5 students
		fmt.Printf("%d. %s (ID: %d) - GPA: %.2f - Status: %s\n",
			i+1, student.Student.Name(), student.Student.ID(), student.OverallGPA, student.Status)
	}
}
