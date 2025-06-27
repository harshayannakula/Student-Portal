package reports

import (
	"encoding/json"
	"fmt"
	"log"
	"oops/main/domain"
	"os"
)

type courseResultData struct {
	StudentID  int     `json:"student_id"`
	CourseID   int     `json:"course_id"`
	CourseName string  `json:"course_name"`
	Grade      string  `json:"grade"`
	Semester   int     `json:"semester"`
	Credits    float64 `json:"credits"`
}

type courseData struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Credits float64 `json:"credits"`
}

func parseGrade(gradeStr string) (domain.AlphabeticGrade, error) {
	gradeMap := map[string]domain.AlphabeticGrade{
		"O":  domain.O,
		"A+": domain.Aplus,
		"A":  domain.A,
		"B+": domain.Bplus,
		"B":  domain.B,
		"C":  domain.C,
		"F":  domain.F,
	}

	if grade, exists := gradeMap[gradeStr]; exists {
		return grade, nil
	}
	return domain.F, fmt.Errorf("invalid grade: %s", gradeStr)
}

func LoadCourses(regis *domain.Registrar) []domain.Course {
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
		regis.AddCourse(course)
	}
	return courseMap
}

func LoadCourseResults() []domain.CourseResult {
	var resultsData []courseResultData
	data, err := os.ReadFile("courseResults.json")
	if err != nil {
		log.Fatal("Failed to read courseResults.json:", err)
	}

	err = json.Unmarshal(data, &resultsData)
	if err != nil {
		log.Fatal("Failed to unmarshal course results:", err)
	}

	var results []domain.CourseResult
	for _, rd := range resultsData {
		grade, err := parseGrade(rd.Grade)
		if err != nil {
			log.Printf("Warning: %v for student ID %d", err, rd.StudentID)
			continue
		}
		courseResult := domain.NewCourseResult(rd.StudentID, rd.CourseID, rd.CourseName, grade, rd.Semester, rd.Credits)
		results = append(results, courseResult)
	}
	return results
}

func GetAtRiskStudents() []domain.StudentRecord {
	var atRisk []domain.StudentRecord
	for _, student := range grg.students {
		if student.Status == "At Risk" {
			atRisk = append(atRisk, student)
		}
	}
	return atRisk
}

func GetDeanListStudents() []domain.StudentRecord {
	var deanList []domain.StudentRecord
	for _, student := range grg.students {
		if student.Status == "Dean's List" {
			deanList = append(deanList, student)
		}
	}
	return deanList
}
