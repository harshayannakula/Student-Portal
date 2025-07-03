package main

import (
	"fmt"
	"student-portal/internal" // ✅ adjust this if your module path is different
)

func main() {
	// Run GPA Histogram Analysis
	var hist map[string]int
	hist, err := internal.GenerateGPAHistogramFromFiles("courseResults.json", "students.json")
	if err != nil {
		fmt.Println("❌ Failed to generate GPA histogram:", err)
	} else {
		//fmt.Println("✅ GPA Histogram generated:", hist)
		err = internal.ExportGPAHistogramChart(hist, "gpa_histogram.png")
		if err != nil {
			fmt.Println("❌ Failed to export histogram chart:", err)
		} else {
			fmt.Println("✅ GPA Trends Chart Generated")
		}
	}
	// Dean List Chart
	if err := internal.ExportDeanListChart("courseResults.json", "students.json", "dean_list.png"); err != nil {
		fmt.Println("❌ Dean List chart export failed:", err)
	} else {
		fmt.Println("✅ Dean List Chart Generated.")
	}

	// At-Risk Students Chart
	if err := internal.ExportAtRiskChart("courseResults.json", "students.json", "at_risk_students.png"); err != nil {
		fmt.Println("❌ At-Risk chart export failed:", err)
	} else {
		fmt.Println("✅ At-Risk Chart Generated.")
	}

	// Placement offer categorization
	offers, err := internal.LoadOffers("placement_offers.json")
	if err != nil {
		fmt.Println("❌ Failed to load offers:", err)
	} else {
		categorized := internal.CategorizeOffers(offers)
		err = internal.ExportCategorizedOffers("placement_chart.json", categorized)
		if err != nil {
			fmt.Println("❌ Export failed:", err)
		}
		//} else {
		//	fmt.Print("✅ Offers categorized and saved to placement_chart.json")
		//}
	}

	// Export placement bar chart
	err = internal.ExportPlacementBarChart("placement_chart.json", "placement_chart.png")
	if err != nil {
		fmt.Println("❌ Placement chart export failed:", err)
	} else {
		fmt.Println("✅ Placement Chart Generated.")
	}

	//Company Wise Selection Metrics
	if err := internal.ExportCompanySelectionChart("placement_offers.json", "company_selection.png", "company_selection.json"); err != nil {
		fmt.Println("❌ Company Selection chart export failed:", err)
	} else {
		fmt.Println("✅ Company Selection Chart Generated.")
	}
}
