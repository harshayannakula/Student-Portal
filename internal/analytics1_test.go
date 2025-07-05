package internal

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGenerateGPAHistogramFromFiles(t *testing.T) {
	//Create dummy students.json
	type testStudent struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
	students := []testStudent{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	}

	//Create Dummy courseResults.json using grade as string
	studentBytes, _ := json.Marshal(students)
	_ = os.WriteFile("test_students.json", studentBytes, 0644)
	defer os.Remove("test_students.json")

	type testCourseResult struct {
		StudentId int     `json:"StudentId"`
		Semester  int     `json:"Semester"`
		CourseId  int     `json:"CourseId"`
		Credits   float64 `json:"Credits"`
		Grade     string  `json:"Grade"`
	}
	courseResults := []testCourseResult{
		{StudentId: 1, Semester: 1, CourseId: 101, Credits: 4.0, Grade: "A"},
		{StudentId: 1, Semester: 1, CourseId: 102, Credits: 3.0, Grade: "B+"},
		{StudentId: 2, Semester: 1, CourseId: 101, Credits: 4.0, Grade: "C"},
	}

	resultBytes, _ := json.Marshal(courseResults)
	_ = os.WriteFile("test_results.json", resultBytes, 0644)
	defer os.Remove("test_results.json")

	//Run Histogram Generator
	hist, err := GenerateGPAHistogramFromFiles("test_results.json", "test_students.json")
	if err != nil {
		t.Fatalf("Error generating histogram: %v", err)
	}

	//Validate that histogram contains expected GPA buckets
	if len(hist) == 0 {
		t.Errorf("Expected GPA buckets, got empty histogram")
	}

	t.Logf("Histogram: %v", hist)
}

func TestExportGPAHistogramChart(t *testing.T) {
	//Sample Histogram Data
	hist := map[string]int{
		"5-5.9": 2,
		"6-6.9": 4,
		"7-7.9": 6,
	}

	//Generating the Chart
	outputFile := "test_gpa_histogram.png"
	err := ExportGPAHistogramChart(hist, outputFile)
	if err != nil {
		t.Fatalf("Error exporting histogram: %v", err)
	}
	defer os.Remove(outputFile)

	//Check if file exists and has content
	info, err := os.Stat(outputFile)
	if os.IsNotExist(err) {
		t.Fatalf("Output file %s does not exist", outputFile)
	}
	if info.Size() < 1000 {
		t.Fatalf("Output file %s is too small Possibly Empty or Corrupt", outputFile)
	}
}
func TestGetGPABucket(t *testing.T) {
	tests := []struct {
		gpa      float64
		expected string
	}{
		{3.9, "<4"},
		{4.5, "4–4.9"},
		{5.5, "5–5.9"},
		{6.5, "6–6.9"},
		{7.5, "7–7.9"},
		{8.5, "8–8.9"},
		{9.1, "9–10"},
	}

	for _, tt := range tests {
		if result := getGPABucket(tt.gpa); result != tt.expected {
			t.Errorf("Expected %s, got %s for GPA %f", tt.expected, result, tt.gpa)
		}
	}
}
func TestExportGPAHistogramChart_Error(t *testing.T) {
	hist := map[string]int{"6-6.9": 2}
	// Intentionally giving bad path
	err := ExportGPAHistogramChart(hist, "/badpath/gpa_chart.png")
	if err == nil {
		t.Errorf("Expected error for bad file path")
	}
}
func TestGenerateGPAHistogramFromFiles_FileNotFound(t *testing.T) {
	_, err := GenerateGPAHistogramFromFiles("nonexistent.json", "students.json")
	if err == nil {
		t.Errorf("Expected error for missing file, got nil")
	}
}

func TestExportGPAHistogramChart_EmptyOrInvalidData(t *testing.T) {
	err := ExportGPAHistogramChart(map[string]int{}, "invalid_chart.png")
	if err != nil {
		t.Errorf("Expected no error for empty chart, got: %v", err)
	}
}

func TestExportGPAHistogramChart_EmptyBuckets(t *testing.T) {
	hist := map[string]int{
		"<4": 0, "4–4.9": 0, "5–5.9": 0, "6–6.9": 0, "7–7.9": 0, "8–8.9": 0, "9–10": 1,
	}
	err := ExportGPAHistogramChart(hist, "test_empty_gpa_chart.png")
	assert.NoError(t, err)
	defer os.Remove("test_empty_gpa_chart.png")
	defer os.Remove("gpa_histogram.json")
}
