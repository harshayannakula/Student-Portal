package internal

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

// Sample data
var sampleOffers = []PlacementOffer{
	{"Google", 30.0, 3, "SWE"},
	{"TCS", 5.0, 10, "ASE"},
	{"Amazon", 20.0, 2, "Cloud Engg"},
	{"Infosys", 6.5, 6, "SE"},
	{"Microsoft", 45.0, 1, "SDE"},
}

func writeTempJSON(t *testing.T, data interface{}) string {
	t.Helper()
	tmpFile := filepath.Join(t.TempDir(), "offers.json")
	content, _ := json.MarshalIndent(data, "", "  ")
	if err := os.WriteFile(tmpFile, content, 0644); err != nil {
		t.Fatalf("Failed to write temp file: %v", err)
	}
	return tmpFile
}

func TestCategorizeOffers(t *testing.T) {
	categorized := CategorizeOffers(sampleOffers)
	if len(categorized["marquee"]) != 3 {
		t.Errorf("Expected 3 marquee offers, got %d", len(categorized["marquee"]))
	}
	if len(categorized["dream"]) != 2 {
		t.Errorf("Expected 2 dream offers, got %d", len(categorized["dream"]))
	}
	if len(categorized["super_dream"]) != 0 {
		t.Errorf("Expected 0 super_dream offer, got %d", len(categorized["super_dream"]))
	}
}

func TestExportCategorizedOffers(t *testing.T) {
	categorized := CategorizeOffers(sampleOffers)
	tmpFile := filepath.Join(t.TempDir(), "categorized.json")
	err := ExportCategorizedOffers(tmpFile, categorized)
	if err != nil {
		t.Fatalf("Failed to export JSON: %v", err)
	}
	// File should exist
	if _, err := os.Stat(tmpFile); err != nil {
		t.Errorf("Exported file does not exist: %v", err)
	}
}

func TestExportPlacementBarChart(t *testing.T) {
	categorized := CategorizeOffers(sampleOffers)
	tmpFile := writeTempJSON(t, categorized)
	output := filepath.Join(t.TempDir(), "placement_chart.png")

	err := ExportPlacementBarChart(tmpFile, output)
	if err != nil {
		t.Fatalf("Bar chart export failed: %v", err)
	}
	if _, err := os.Stat(output); err != nil {
		t.Errorf("Expected PNG chart not found: %v", err)
	}
}

func TestExportCompanySelectionChart(t *testing.T) {
	tmpFile := writeTempJSON(t, sampleOffers)
	outputPNG := filepath.Join(t.TempDir(), "company_selection.png")
	outputJSON := filepath.Join(t.TempDir(), "company_selection.json")

	err := ExportCompanySelectionChart(tmpFile, outputPNG, outputJSON)
	if err != nil {
		t.Fatalf("Company selection chart failed: %v", err)
	}

	if _, err := os.Stat(outputPNG); err != nil {
		t.Errorf("Expected PNG chart missing: %v", err)
	}
	if _, err := os.Stat(outputJSON); err != nil {
		t.Errorf("Expected JSON chart data missing: %v", err)
	}
}
func TestLoadOffers(t *testing.T) {
	validJSON := `[{"CompanyName":"X","PackageLPA":10.0,"NumStudents":5,"JobTitle":"Engineer"}]`
	tmp := "test_offers.json"
	_ = os.WriteFile(tmp, []byte(validJSON), 0644)
	defer os.Remove(tmp)

	offers, err := LoadOffers(tmp)
	if err != nil || len(offers) != 1 {
		t.Errorf("Expected 1 offer, got %v, err: %v", offers, err)
	}

	_, err = LoadOffers("nonexistent.json")
	if err == nil {
		t.Errorf("Expected error for missing file")
	}
}
func TestGetCategory(t *testing.T) {
	if cat := getCategory(5.0); cat != "dream" {
		t.Errorf("Expected 'dream' category for package < 10, got %s", cat)
	}
}
func TestExportCategorizedOffers_Error(t *testing.T) {
	data := map[string][]PlacementOffer{"dream": {
		{CompanyName: "Z", PackageLPA: 5.0, NumStudents: 1, JobTitle: "Intern"},
	}}
	err := ExportCategorizedOffers("/badpath/offers.json", data)
	if err == nil {
		t.Errorf("Expected error when writing to an invalid path")
	}
}
func TestLoadOffers_FileNotFound(t *testing.T) {
	_, err := LoadOffers("missing.json")
	if err == nil {
		t.Errorf("Expected error for missing file")
	}
}

func TestExportCategorizedOffers_InvalidPath(t *testing.T) {
	offers := map[string][]PlacementOffer{}
	err := ExportCategorizedOffers("/invalid_path/file.json", offers)
	if err == nil {
		t.Errorf("Expected write error for invalid path")
	}
}

func TestGetCategory_SuperDreamCase(t *testing.T) {
	category := getCategory(15.0)
	if category != "super_dream" {
		t.Errorf("Expected 'super_dream', got %s", category)
	}
}
