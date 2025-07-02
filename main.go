package main

import (
	"fmt"
	"student-portal/internal" // ✅ adjust this if your module name is different
)

func main() {
	// Run GPA Histogram Analysis
	hist, err := internal.GenerateGPAHistogramFromFiles("courseResults.json", "students.json")
	if err != nil {
		fmt.Println("❌ Failed to generate GPA histogram:", err)
	} else {
		fmt.Println("✅ GPA Histogram generated:", hist)
		err = internal.ExportGPAHistogramChart(hist, "gpa_histogram.png")
		if err != nil {
			fmt.Println("❌ Failed to export histogram chart:", err)
		} else {
			fmt.Println("✅ Chart saved to gpa_histogram.png")
		}
	}

}
