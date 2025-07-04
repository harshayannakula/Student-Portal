package internal

import (
	"encoding/json"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
	"os"
	"sort"
	"strings"
)

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
