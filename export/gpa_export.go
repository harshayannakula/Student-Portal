package export

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"example.com/m/v3/domain"
)

func ExportDetailedTranscript(path string, students []domain.StudentRecord) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	header := []string{"Student_ID", "Student_Name", "Semester", "Semester_GPA", "Overall_GPA", "Status"}
	if err := w.Write(header); err != nil {
		return err
	}

	for _, student := range students {
		for _, semester := range student.Semesters {
			record := []string{
				strconv.Itoa(semester.Student.ID()),
				semester.Student.Name(),
				strconv.Itoa(semester.Semester),
				fmt.Sprintf("%.2f", semester.Gpa),
				fmt.Sprintf("%.2f", student.OverallGPA),
				student.Status,
			}
			if err := w.Write(record); err != nil {
				return err
			}
		}
	}

	return nil
}

func ExportSummaryReport(path string, students []domain.StudentRecord) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	header := []string{"Student_ID", "Student_Name", "Total_Semesters", "Overall_GPA", "Status"}
	if err := w.Write(header); err != nil {
		return err
	}

	for _, student := range students {
		record := []string{
			strconv.Itoa(student.Student.ID()),
			student.Student.Name(),
			strconv.Itoa(len(student.Semesters)),
			fmt.Sprintf("%.2f", student.OverallGPA),
			student.Status,
		}
		if err := w.Write(record); err != nil {
			return err
		}
	}

	return nil
}

func ExportAtRiskStudents(path string, students []domain.StudentRecord) error {
	var atRiskStudents []domain.StudentRecord
	for _, student := range students {
		if student.Status == "At Risk" {
			atRiskStudents = append(atRiskStudents, student)
		}
	}
	return ExportSummaryReport(path, atRiskStudents)
}

func ExportDeanListStudents(path string, students []domain.StudentRecord) error {
	var deanListStudents []domain.StudentRecord
	for _, student := range students {
		if student.Status == "Dean's List" {
			deanListStudents = append(deanListStudents, student)
		}
	}
	return ExportSummaryReport(path, deanListStudents)
}

func ExportStatistics(path string, stats map[string]interface{}) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	header := []string{"Metric", "Value"}
	if err := w.Write(header); err != nil {
		return err
	}

	records := [][]string{
		{"Total Students", fmt.Sprintf("%v", stats["total_students"])},
		{"At Risk Count", fmt.Sprintf("%v", stats["at_risk_count"])},
		{"Dean's List Count", fmt.Sprintf("%v", stats["dean_list_count"])},
		{"Normal Count", fmt.Sprintf("%v", stats["normal_count"])},
		{"Average GPA", fmt.Sprintf("%.2f", stats["average_gpa"])},
		{"At Risk Percentage", fmt.Sprintf("%.1f%%", stats["at_risk_percentage"])},
		{"Dean's List Percentage", fmt.Sprintf("%.1f%%", stats["dean_list_percentage"])},
	}

	for _, record := range records {
		if err := w.Write(record); err != nil {
			return err
		}
	}

	return nil
}
