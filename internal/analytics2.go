package internal

import (
	"encoding/json"
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"os"
	"sort"
)

// ExportDeanListChart plots students with GPA > 6
func ExportDeanListChart(courseResultsFile, studentsFile, outputFile string) error {
	return exportFilteredGPAChart(courseResultsFile, studentsFile, outputFile, func(gpa float64) bool {
		return gpa > 6.0
	}, "Dean's List (GPA > 6)")
}

// ExportAtRiskChart plots students with GPA < 5
func ExportAtRiskChart(courseResultsFile, studentsFile, outputFile string) error {
	return exportFilteredGPAChart(courseResultsFile, studentsFile, outputFile, func(gpa float64) bool {
		return gpa < 5.0
	}, "At-Risk Students (GPA < 5)")
}

// Helper function to generate student name vs GPA bar chart
func exportFilteredGPAChart(courseResultsFile, studentsFile, outputFile string, filter func(float64) bool, title string) error {
	// Load course results
	var courseResults []CourseResult
	cData, err := os.ReadFile(courseResultsFile)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(cData, &courseResults); err != nil {
		return err
	}

	// Load student names
	type studentData struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
	var studentRaw []studentData
	sData, err := os.ReadFile(studentsFile)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(sData, &studentRaw); err != nil {
		return err
	}
	students := map[int]string{}
	for _, s := range studentRaw {
		students[s.ID] = s.Name
	}

	// Step 1: Accumulate grade points per student
	type agg struct {
		totalPoints  float64
		totalCredits float64
	}
	scoreMap := map[int]*agg{}

	gradePoints := map[string]float64{
		"O":  10.0,
		"A+": 9.0,
		"A":  8.0,
		"B+": 7.0,
		"B":  6.0,
		"C":  5.0,
		"F":  0.0,
	}

	for _, cr := range courseResults {
		gp, ok := gradePoints[cr.Grade.String()]
		if !ok {
			continue
		}
		if _, exists := scoreMap[cr.StudentId]; !exists {
			scoreMap[cr.StudentId] = &agg{}
		}
		scoreMap[cr.StudentId].totalPoints += gp * cr.Credits
		scoreMap[cr.StudentId].totalCredits += cr.Credits
	}

	// Step 2: Apply filter and collect data
	type record struct {
		Name string
		GPA  float64
	}
	var selected []record
	for id, data := range scoreMap {
		if data.totalCredits == 0 {
			continue
		}
		gpa := data.totalPoints / data.totalCredits
		if filter(gpa) {
			selected = append(selected, record{
				Name: students[id],
				GPA:  gpa,
			})
		}
	}

	// Sort by GPA descending
	sort.Slice(selected, func(i, j int) bool {
		return selected[i].GPA > selected[j].GPA
	})

	// Create chart
	p := plot.New()
	p.Title.Text = title
	p.X.Label.Text = "Students"
	p.Y.Label.Text = "GPA"
	p.Title.TextStyle.Font.Size = vg.Points(14)
	p.X.Label.TextStyle.Font.Size = vg.Points(12)
	p.Y.Label.TextStyle.Font.Size = vg.Points(12)
	p.X.Tick.Label.Rotation = 0
	p.X.Tick.Label.Font.Size = vg.Points(10)
	p.Add(plotter.NewGrid())

	labels := make([]string, len(selected))
	values := make(plotter.Values, len(selected))
	for i, s := range selected {
		labels[i] = s.Name
		values[i] = s.GPA
	}

	if len(selected) == 0 {
		return fmt.Errorf("plotter: no data points (no students matched filter)")
	}

	bars, err := plotter.NewBarChart(values, vg.Points(25))
	if err != nil {
		return err
	}
	bars.Color = plotutil.Color(3)
	bars.LineStyle.Width = 0
	bars.Offset = vg.Points(3)
	p.Add(bars)
	p.NominalX(labels...)

	// Save JSON (bar chart data)
	export := map[string]float64{}
	for _, s := range selected {
		export[s.Name] = s.GPA
	}
	jsonName := outputFile[:len(outputFile)-4] + ".json"

	data, err := json.MarshalIndent(export, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(jsonName, data, 0644); err != nil {
		return err
	}

	return p.Save(14*vg.Inch, 5*vg.Inch, outputFile)
}
