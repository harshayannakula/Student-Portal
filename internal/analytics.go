package internal

import (
	"encoding/json"
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"image/color"
	"os"
	"sort"
	"strings"
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

	// Build academic records
	records := map[int]*AcademicRecord{}
	for _, cr := range courseResults {
		if _, ok := records[cr.StudentId]; !ok {
			records[cr.StudentId] = NewAcademicRecord(cr.StudentId)
		}
		records[cr.StudentId].AddResult(cr, cr.Semester)
	}

	// Filter and collect eligible students
	type record struct {
		Name string
		GPA  float64
	}
	var selected []record
	for id, rec := range records {
		if filter(rec.CGPA) {
			selected = append(selected, record{
				Name: students[id],
				GPA:  rec.CGPA,
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

// PlacementOffer represents a placement offer record
type PlacementOffer struct {
	CompanyName string  `json:"CompanyName"`
	PackageLPA  float64 `json:"PackageLPA"`
	NumStudents int     `json:"NumStudents"`
	JobTitle    string  `json:"JobTitle"`
}

func getCategory(packageLPA float64) string {
	switch {
	case packageLPA >= 20:
		return "marquee"
	case packageLPA >= 10:
		return "super_dream"
	default:
		return "dream"
	}
}

func LoadOffers(filename string) ([]PlacementOffer, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var offers []PlacementOffer
	if err := json.Unmarshal(bytes, &offers); err != nil {
		return nil, err
	}
	return offers, nil
}

func CategorizeOffers(offers []PlacementOffer) map[string][]PlacementOffer {
	result := map[string][]PlacementOffer{
		"marquee":     {},
		"super_dream": {},
		"dream":       {},
	}
	for _, offer := range offers {
		category := getCategory(offer.PackageLPA)
		result[category] = append(result[category], offer)
	}
	return result
}

func ExportCategorizedOffers(filename string, categorizedOffers map[string][]PlacementOffer) error {
	data, err := json.MarshalIndent(categorizedOffers, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func ExportPlacementBarChart(inputFile, outputFile string) error {
	type Offer struct {
		CompanyName string  `json:"CompanyName"`
		PackageLPA  float64 `json:"PackageLPA"`
		NumStudents int     `json:"NumStudents"`
		JobTitle    string  `json:"JobTitle"`
	}

	// Load JSON
	var categorized map[string][]Offer
	data, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, &categorized); err != nil {
		return err
	}

	// Aggregate: map[company][category]count
	type CompStat struct {
		Company string
		Counts  map[string]int
	}
	companiesMap := map[string]*CompStat{}
	categories := []string{"dream", "super_dream", "marquee"}

	for cat, offers := range categorized {
		for _, offer := range offers {
			if _, ok := companiesMap[offer.CompanyName]; !ok {
				companiesMap[offer.CompanyName] = &CompStat{
					Company: offer.CompanyName,
					Counts:  map[string]int{"dream": 0, "super_dream": 0, "marquee": 0},
				}
			}
			companiesMap[offer.CompanyName].Counts[cat] += offer.NumStudents
		}
	}

	// Sort companies
	var companyNames []string
	for name := range companiesMap {
		companyNames = append(companyNames, name)
	}
	sort.Strings(companyNames)

	// Prepare plot
	p := plot.New()
	p.Title.Text = "Placement by Company and Category"
	p.X.Label.Text = "Company"
	p.Y.Label.Text = "Number of Students"
	p.X.Label.TextStyle.Font.Size = vg.Points(12)
	p.Y.Label.TextStyle.Font.Size = vg.Points(12)
	p.Title.TextStyle.Font.Size = vg.Points(14)
	p.X.Tick.Label.Rotation = 0
	p.X.Tick.Label.Font.Size = vg.Points(10)

	//groupWidth := vg.Points(30)
	barWidth := vg.Points(8)
	categoryColors := map[string]color.Color{
		"dream":       color.RGBA{R: 102, G: 187, B: 106, A: 255}, // green
		"super_dream": color.RGBA{R: 51, G: 153, B: 255, A: 255},  // blue
		"marquee":     color.RGBA{R: 255, G: 102, B: 102, A: 255}, // red
	}

	// Each category will have a separate bar plot slice
	legend := map[string]*plotter.BarChart{}
	for i, cat := range categories {
		values := make(plotter.Values, len(companyNames))
		for j, name := range companyNames {
			values[j] = float64(companiesMap[name].Counts[cat])
		}
		bars, err := plotter.NewBarChart(values, barWidth)
		if err != nil {
			return err
		}
		bars.Color = categoryColors[cat]
		bars.Offset = barWidth * vg.Length(i-1) // center offset: -1, 0, +1
		p.Add(bars)
		p.Legend.Add(strings.Title(strings.ReplaceAll(cat, "_", " ")), bars)
		legend[cat] = bars
	}

	// Set ticks
	p.NominalX(companyNames...)

	p.Legend.Top = true
	p.Legend.XOffs = -vg.Points(10)
	p.Legend.YOffs = vg.Points(10)
	p.Legend.TextStyle.Font.Size = vg.Points(10)

	// Save chart
	return p.Save(14*vg.Inch, 5*vg.Inch, outputFile)
}

func ExportCompanySelectionChart(inputFile, outputImage, outputJSON string) error {
	type Offer struct {
		CompanyName string  `json:"CompanyName"`
		PackageLPA  float64 `json:"PackageLPA"`
		NumStudents int     `json:"NumStudents"`
		JobTitle    string  `json:"JobTitle"`
	}

	// Load JSON data
	var offers []Offer
	data, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, &offers); err != nil {
		return err
	}

	// Aggregate: map[company] = {totalStudents, totalPackage}
	type CompanyStat struct {
		TotalStudents int
		TotalPackage  float64
	}
	stats := map[string]CompanyStat{}
	for _, offer := range offers {
		stat := stats[offer.CompanyName]
		stat.TotalStudents += offer.NumStudents
		stat.TotalPackage += offer.PackageLPA * float64(offer.NumStudents)
		stats[offer.CompanyName] = stat
	}

	// Prepare sorted keys and values
	var companies []string
	for name := range stats {
		companies = append(companies, name)
	}
	sort.Strings(companies)

	nums := make(plotter.Values, len(companies))
	avgs := make([]float64, len(companies))
	exportData := map[string]map[string]float64{}

	for i, name := range companies {
		stat := stats[name]
		nums[i] = float64(stat.TotalStudents)
		avg := stat.TotalPackage / float64(stat.TotalStudents)
		avgs[i] = avg
		exportData[name] = map[string]float64{
			"num_students": float64(stat.TotalStudents),
			"avg_package":  avg,
		}
	}

	// Plotting
	p := plot.New()
	p.Title.Text = "Company-wise Selection Metrics"
	p.X.Label.Text = "Company"
	p.Y.Label.Text = "Number of Students"

	barChart, err := plotter.NewBarChart(nums, vg.Points(25))
	if err != nil {
		return err
	}
	barChart.LineStyle.Width = 0
	barChart.Color = color.RGBA{R: 100, G: 180, B: 255, A: 255}
	barChart.LineStyle.Width = 0
	barChart.Offset = vg.Points(0)
	p.Add(barChart)
	p.Legend.Add("Number of Students", barChart)

	// Add right Y-axis
	rightAxis := plot.New()
	rightAxis.Y.Label.Text = "Average Package (LPA)"
	rightAxis.Y.Tick.Label.Color = color.RGBA{R: 255, G: 80, B: 80, A: 255}

	// Plot line for average package
	linePoints := make(plotter.XYs, len(companies))
	for i := range companies {
		linePoints[i].X = float64(i)
		linePoints[i].Y = avgs[i]
	}
	line, err := plotter.NewLine(linePoints)
	if err != nil {
		return err
	}
	line.Color = color.RGBA{R: 255, G: 80, B: 80, A: 255}
	line.Width = vg.Points(2)
	p.Add(line)
	p.Legend.Add("Average Package (LPA)", line)
	p.NominalX(companies...)

	line.Color = color.RGBA{R: 255, G: 80, B: 80, A: 255}
	p.Add(line)
	p.Legend.Top = true
	p.Legend.XOffs = -vg.Points(10)
	p.Legend.YOffs = vg.Points(10)
	p.Legend.TextStyle.Font.Size = vg.Points(10)

	// Save image
	if err := p.Save(14*vg.Inch, 5*vg.Inch, outputImage); err != nil {
		return err
	}

	// Save JSON
	jsonBytes, err := json.MarshalIndent(exportData, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(outputJSON, jsonBytes, 0644)
}
