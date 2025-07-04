package internal

import (
	"encoding/json"
	"os"
	"testing"
)

func TestExportDeanAndAtRiskCharts(t *testing.T) {
	// --- Step 1: Create mock student data ---
	students := []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}{
		{1, "Alice"},
		{2, "Bob"},
		{3, "Charlie"},
	}
	sData, _ := json.Marshal(students)
	_ = os.WriteFile("test_students.json", sData, 0644)
	defer os.Remove("test_students.json")

	// --- Step 2: Create mock course result data ---
	type TestCourseResult struct {
		StudentId  int     `json:"student_id"`
		Semester   int     `json:"semester"`
		CourseId   int     `json:"course_id"`
		CourseName string  `json:"course_name"`
		Credits    float64 `json:"credits"`
		Grade      string  `json:"grade"` // <-- important: string
	}

	courses := []TestCourseResult{
		{StudentId: 1, Semester: 1, CourseId: 1, Credits: 3.0, Grade: "A+"},
		{StudentId: 2, Semester: 1, CourseId: 1, Credits: 3.0, Grade: "F"},
	}

	rData, _ := json.Marshal(courses)
	_ = os.WriteFile("test_results.json", rData, 0644)
	defer os.Remove("test_results.json")

	// --- Step 3: Test Dean List Chart (GPA > 6) ---
	deanPNG := "test_dean_chart.png"
	deanJSON := "test_dean_chart.json"
	err := ExportDeanListChart("test_results.json", "test_students.json", deanPNG)
	if err != nil {
		t.Errorf("Dean list chart failed: %v", err)
	}
	checkFile(t, deanPNG)
	checkFile(t, deanJSON)

	// --- Step 4: Test At-Risk Chart (GPA < 5) ---
	riskPNG := "test_risk_chart.png"
	riskJSON := "test_risk_chart.json"
	err = ExportAtRiskChart("test_results.json", "test_students.json", riskPNG)
	if err != nil {
		t.Errorf("At-risk chart failed: %v", err)
	}
	checkFile(t, riskPNG)
	checkFile(t, riskJSON)

	// Cleanup
	defer os.Remove(deanPNG)
	defer os.Remove(deanJSON)
	defer os.Remove(riskPNG)
	defer os.Remove(riskJSON)
}

// checkFile verifies file exists and is not empty
func checkFile(t *testing.T, filename string) {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		t.Errorf("Expected file %s does not exist", filename)
		return
	}
	if info.Size() == 0 {
		t.Errorf("File %s is empty", filename)
	}
}
