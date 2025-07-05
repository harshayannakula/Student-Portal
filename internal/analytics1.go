package internal

import (
	"encoding/json"
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"os"
)

func GenerateGPAHistogramFromFiles(courseResultsFile, studentsFile string) (map[string]int, error) {
	// Step 1: Read and parse courseResults.json
	var courseResults []CourseResult
	cData, err := os.ReadFile(courseResultsFile)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(cData, &courseResults); err != nil {
		return nil, err
	}

	// Step 2: Read and parse students.json
	type studentData struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
	var studentRaw []studentData
	sData, err := os.ReadFile(studentsFile)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(sData, &studentRaw); err != nil {
		return nil, err
	}

	// Step 3: Build student map for name lookup
	students := map[int]Student{}
	for _, s := range studentRaw {
		students[s.ID] = NewStudent(s.ID, s.Name)
	}

	// Step 4: Build academic records
	records := map[int]*AcademicRecord{}
	for _, cr := range courseResults {
		if _, exists := records[cr.StudentId]; !exists {
			records[cr.StudentId] = NewAcademicRecord(cr.StudentId)
		}
		records[cr.StudentId].AddResult(cr, cr.Semester)
	}

	// Step 5: Set Status and attach Name
	for _, record := range records {
		record.Status = NewGPACalculator().DetermineStatus(record.CGPA)
		// optional: assign Name if needed later
	}

	// Step 6: Create histogram
	hist := map[string]int{}
	for _, record := range records {
		bucket := getGPABucket(record.CGPA)
		hist[bucket]++
	}

	return hist, nil
}
func getGPABucket(gpa float64) string {
	switch {
	case gpa < 4.0:
		return "<4"
	case gpa < 5.0:
		return "4–4.9"
	case gpa < 6.0:
		return "5–5.9"
	case gpa < 7.0:
		return "6–6.9"
	case gpa < 8.0:
		return "7–7.9"
	case gpa < 9.0:
		return "8–8.9"
	default:
		return "9–10"
	}
}
func ExportGPAHistogramChart(hist map[string]int, filename string) error {
	p := plot.New()
	p.Title.Text = "GPA Distribution Histogram"
	p.Title.TextStyle.Font.Size = vg.Points(14)
	p.X.Label.Text = "GPA Range"
	p.Y.Label.Text = "Number of Students"
	p.X.Label.TextStyle.Font.Size = vg.Points(12)
	p.Y.Label.TextStyle.Font.Size = vg.Points(12)
	p.Add(plotter.NewGrid())

	// Predefined ordered GPA bucket labels
	labels := []string{"<4", "4–4.9", "5–5.9", "6–6.9", "7–7.9", "8–8.9", "9–10"}
	values := make(plotter.Values, len(labels))
	for i, label := range labels {
		values[i] = float64(hist[label]) // defaults to 0 if missing
	}

	// Draw bar chart
	bar, err := plotter.NewBarChart(values, vg.Points(30))
	if err != nil {
		return err
	}
	bar.LineStyle.Width = 0
	bar.Color = plotutil.Color(1)
	bar.Offset = vg.Points(5)
	p.Add(bar)
	p.NominalX(labels...)

	// Add value labels above each bar
	for i, v := range values {
		if v > 0 {
			lbl, err := plotter.NewLabels(plotter.XYLabels{
				XYs:    []plotter.XY{{X: float64(i), Y: v + 0.3}},
				Labels: []string{fmt.Sprintf("%.0f", v)},
			})
			if err != nil {
				return err
			}
			p.Add(lbl)
		}
	}
	// Save JSON as well
	jsonData, err := json.MarshalIndent(hist, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile("gpa_histogram.json", jsonData, 0644)
	if err != nil {
		return err
	}

	return p.Save(10*vg.Inch, 4*vg.Inch, filename)
}
