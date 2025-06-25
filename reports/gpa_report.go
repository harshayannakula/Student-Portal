package reports

import (
	"fmt"
	"math/rand"
	"time"

	"example.com/m/v3/domain"
)

type GPAReportGenerator struct {
	calculator *domain.GPACalculator
	students   []domain.StudentRecord
}

func NewGPAReportGenerator() *GPAReportGenerator {
	return &GPAReportGenerator{
		calculator: domain.NewGPACalculator(),
		students:   make([]domain.StudentRecord, 0),
	}
}

func (grg *GPAReportGenerator) GenerateStudentData() {
	rand.Seed(time.Now().UnixNano())

	firstNames := []string{
		"Alice", "Bob", "Charlie", "Diana", "Edward", "Fiona", "George", "Hannah",
		"Ivan", "Julia", "Kevin", "Luna", "Michael", "Nina", "Oscar", "Priya",
		"Quinn", "Rachel", "Steve", "Tina", "Uma", "Victor", "Wendy", "Xavier",
		"Yara", "Zoe", "Adam", "Beth", "Carl", "Dora", "Eric", "Faith",
	}

	lastNames := []string{
		"Anderson", "Brown", "Clark", "Davis", "Evans", "Fisher", "Garcia", "Harris",
		"Jackson", "Kumar", "Lee", "Miller", "Nelson", "O'Connor", "Patel", "Quinn",
		"Rodriguez", "Smith", "Taylor", "Underwood", "Vasquez", "Wilson", "Young", "Zhang",
	}

	for i := 1; i <= 100; i++ {

		firstName := firstNames[rand.Intn(len(firstNames))]
		lastName := lastNames[rand.Intn(len(lastNames))]
		fullName := fmt.Sprintf("%s %s", firstName, lastName)

		student := domain.NewStudent(i, fullName)

		numSemesters := rand.Intn(8) + 1
		semesters := make([]domain.StudentGPA, numSemesters)

		for j := 1; j <= numSemesters; j++ {
			gpa := generateRealisticGPA()
			semesters[j-1] = domain.StudentGPA{
				Student:  student,
				Semester: j,
				Gpa:      gpa,
			}
		}

		overallGPA := grg.calculator.CalculateOverallGPA(semesters)
		status := grg.calculator.DetermineStatus(overallGPA)

		studentRecord := domain.StudentRecord{
			Student:    student,
			Semesters:  semesters,
			OverallGPA: overallGPA,
			Status:     status,
		}

		grg.students = append(grg.students, studentRecord)
	}
}

func generateRealisticGPA() float64 {

	randVal := rand.Float64()

	if randVal < 0.15 {

		return rand.Float64() * 4
	} else if randVal < 0.85 {

		return 4 + rand.Float64()*4
	} else {

		return 8 + rand.Float64()*2
	}
}

func (grg *GPAReportGenerator) GetStudents() []domain.StudentRecord {
	return grg.students
}

func (grg *GPAReportGenerator) GetAtRiskStudents() []domain.StudentRecord {
	var atRisk []domain.StudentRecord
	for _, student := range grg.students {
		if student.Status == "At Risk" {
			atRisk = append(atRisk, student)
		}
	}
	return atRisk
}

func (grg *GPAReportGenerator) GetDeanListStudents() []domain.StudentRecord {
	var deanList []domain.StudentRecord
	for _, student := range grg.students {
		if student.Status == "Dean's List" {
			deanList = append(deanList, student)
		}
	}
	return deanList
}

func (grg *GPAReportGenerator) GenerateStatistics() map[string]interface{} {
	total := len(grg.students)
	atRisk := len(grg.GetAtRiskStudents())
	deanList := len(grg.GetDeanListStudents())
	normal := total - atRisk - deanList

	var totalGPA float64
	for _, student := range grg.students {
		totalGPA += student.OverallGPA
	}
	avgGPA := totalGPA / float64(total)

	return map[string]interface{}{
		"total_students":       total,
		"at_risk_count":        atRisk,
		"dean_list_count":      deanList,
		"normal_count":         normal,
		"average_gpa":          avgGPA,
		"at_risk_percentage":   float64(atRisk) / float64(total) * 100,
		"dean_list_percentage": float64(deanList) / float64(total) * 100,
	}
}
