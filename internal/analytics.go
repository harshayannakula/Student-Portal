package internal

import (
	"encoding/json"
	"os"
	"sort"
)

type StudentGPATrend struct {
	StudentID int     `json:"student_id"`
	Name      string  `json:"name"`
	Semester  int     `json:"semester"`
	CGPA      float64 `json:"cgpa"`
}

type CourseResult struct {
	StudentID  int     `json:"student_id"`
	CourseID   int     `json:"course_id"`
	CourseName string  `json:"course_name"`
	Grade      string  `json:"grade"`
	Semester   int     `json:"semester"`
	Credits    float64 `json:"credits"`
}

type Student struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Mapping letter grades to points
var gradePointMap = map[string]float64{
	"O": 10, "A+": 9, "A": 8, "B+": 7, "B": 6, "C": 5, "F": 0,
}

func CalculateGPATrends(courseResults []CourseResult, students []Student) []StudentGPATrend {
	studentNameMap := map[int]string{}
	for _, s := range students {
		studentNameMap[s.ID] = s.Name
	}

	// studentID -> semester -> (sum of gradePoints, total credits)
	temp := map[int]map[int][2]float64{}

	for _, cr := range courseResults {
		points, ok := gradePointMap[cr.Grade]
		if !ok {
			continue // skip unknown grades
		}
		if _, ok := temp[cr.StudentID]; !ok {
			temp[cr.StudentID] = map[int][2]float64{}
		}
		acc := temp[cr.StudentID][cr.Semester]
		acc[0] += points * cr.Credits // total grade points
		acc[1] += cr.Credits          // total credits
		temp[cr.StudentID][cr.Semester] = acc
	}

	// Convert to slice of StudentGPATrend
	var trends []StudentGPATrend
	for sid, semesters := range temp {
		for sem, gp := range semesters {
			cgpa := 0.0
			if gp[1] > 0 {
				cgpa = gp[0] / gp[1]
			}
			trends = append(trends, StudentGPATrend{
				StudentID: sid,
				Name:      studentNameMap[sid],
				Semester:  sem,
				CGPA:      cgpa,
			})
		}
	}

	sort.Slice(trends, func(i, j int) bool {
		if trends[i].StudentID == trends[j].StudentID {
			return trends[i].Semester < trends[j].Semester
		}
		return trends[i].StudentID < trends[j].StudentID
	})

	return trends
}

func ExportGPATrendsToJSON(filename string, data []StudentGPATrend) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}
