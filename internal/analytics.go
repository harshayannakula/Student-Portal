package internal

import (
	"encoding/json"
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"os"
	"strconv"
)

func GenerateGPAHistogramFromFiles(courseResultsFile, studentsFile string) (map[int]int, error) {
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
	hist := map[int]int{}
	for _, record := range records {
		bucket := int(record.CGPA)
		hist[bucket]++
	}

	return hist, nil
}
func ExportGPAHistogramChart(hist map[int]int, filename string) error {
	p := plot.New()
	p.Title.Text = "GPA Distribution Histogram"
	p.Title.TextStyle.Font.Size = vg.Points(14)
	p.X.Label.Text = "GPA Bucket (e.g., 6 = GPA between 6.0–6.9)"
	p.Y.Label.Text = "Number of Students"
	p.X.Label.TextStyle.Font.Size = vg.Points(12)
	p.Y.Label.TextStyle.Font.Size = vg.Points(12)
	p.Add(plotter.NewGrid())

	// Create GPA buckets 0 to 10
	values := make(plotter.Values, 11)
	labels := make([]string, 11)
	for i := 0; i <= 10; i++ {
		values[i] = float64(hist[i]) // missing keys default to 0
		labels[i] = strconv.Itoa(i)
	}

	// Draw bar chart
	bar, err := plotter.NewBarChart(values, vg.Points(25))
	if err != nil {
		return err
	}
	bar.LineStyle.Width = 0
	bar.Color = plotutil.Color(2)
	bar.Offset = vg.Points(5)
	p.Add(bar)
	p.NominalX(labels...)

	// Add labels on top of bars
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

	return p.Save(10*vg.Inch, 4*vg.Inch, filename)
}
