package internal

import (
	"testing"
	"time"
)

func TestNewEligibility(t *testing.T) {
	el := NewEligibility(7.5)
	if el.Requirement() != 7.5 {
		t.Error("NewEligibility or Requirement failed")
	}
}

func TestEligibility_ChangeRequirement(t *testing.T) {
	el := NewEligibility(6.0)
	el.ChangeRequirement(8.0)
	if el.Requirement() != 8.0 {
		t.Error("ChangeRequirement did not update requirement")
	}
}

func TestDrive_GettersSetters(t *testing.T) {
	start := time.Now()
	end := start.Add(24 * time.Hour)
	dr := NewDrive(start, end, "SWE", 8.0, 100, Dream)
	if dr.StartDate() != start || dr.EndDate() != end || dr.RoleName() != "SWE" || dr.CTC() != 100 || dr.JobCategory() != Dream {
		t.Error("Drive getters failed")
	}
	dr.SetStartDate(end)
	dr.SetEndDate(start)
	dr.SetRoleName("DevOps")
	dr.SetCTC(200)
	dr.SetJobCategory(Marquee)
	if dr.StartDate() != end || dr.EndDate() != start || dr.RoleName() != "DevOps" || dr.CTC() != 200 || dr.JobCategory() != Marquee {
		t.Error("Drive setters failed")
	}
}

func TestDrive_Eligibility(t *testing.T) {
	dr := NewDrive(time.Now(), time.Now(), "QA", 8.0, 50, Day)
	a := NewApplicant(Student{id: 1, name: "Bob"}, AcademicRecord{CGPA: 8.1})
	if !dr.Eligibility().checkEligibility(a) {
		t.Error("Eligibility check should pass")
	}
	a2 := NewApplicant(Student{id: 2}, AcademicRecord{CGPA: 7.0})
	if dr.Eligibility().checkEligibility(a2) {
		t.Error("Eligibility check should fail")
	}
}

func TestEligibility_checkEligibility(t *testing.T) {
	tests := []struct {
		name          string
		requirement   float64
		applicantCGPA float64
		expected      bool
	}{
		{"eligible - exact match", 3.5, 3.5, true},
		{"eligible - above requirement", 3.0, 3.8, true},
		{"not eligible - below requirement", 3.5, 3.2, false},
		{"edge case - zero requirement", 0.0, 2.5, true},
		{"edge case - perfect GPA", 3.0, 4.0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			eligibility := &Eligibility{requirement: tt.requirement}
			applicant := &Applicant{AcademicRecord: AcademicRecord{CGPA: tt.applicantCGPA}}

			result := eligibility.checkEligibility(applicant)
			if result != tt.expected {
				t.Errorf("checkEligibility() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestDrive_Applications(t *testing.T) {
	dr := NewDrive(time.Now(), time.Now(), "Analyst", 7.0, 80, Dream)
	app := &Application{id: 1, Applicant: &Applicant{Student: Student{id: 1}}, status: Applied}
	dr.AppendApplication(app)
	if len(dr.Applications()) != 1 {
		t.Error("AppendApplication failed")
	}
}

func TestDrive_HasApplied(t *testing.T) {
	dr := NewDrive(time.Now(), time.Now(), "QA", 8.0, 50, Day)
	app := &Application{id: 1, Applicant: &Applicant{Student: Student{id: 1}}, status: Applied}
	dr.AppendApplication(app)
	if !dr.HasApplied(1) {
		t.Error("HasApplied should return true")
	}
	if dr.HasApplied(2) {
		t.Error("HasApplied should return false")
	}
}

func TestDrive_GetApplicationByID(t *testing.T) {
	dr := NewDrive(time.Now(), time.Now(), "QA", 8.0, 50, Day)
	app := &Application{id: 2, Applicant: &Applicant{Student: Student{id: 2}}, status: Applied}
	dr.AppendApplication(app)
	got, err := dr.GetApplicationByID(2)
	if err != nil || got != app {
		t.Error("GetApplicationByID failed")
	}
	_, err2 := dr.GetApplicationByID(999)
	if err2 == nil {
		t.Error("GetApplicationByID should fail for missing id")
	}
}

func TestDrive_getSelectedApplications(t *testing.T) {
	dr := NewDrive(time.Now(), time.Now(), "QA", 8.0, 50, Day)
	app1 := &Application{id: 1, status: Selected}
	app2 := &Application{id: 2, status: Applied}
	dr.AppendApplication(app1)
	dr.AppendApplication(app2)
	sel := dr.getSelectedApplications()
	if len(sel) != 1 || sel[0] != app1 {
		t.Error("getSelectedApplications failed")
	}
}

func TestDrive_getShortlistedApplications(t *testing.T) {
	dr := NewDrive(time.Now(), time.Now(), "QA", 8.0, 50, Day)
	app1 := &Application{id: 1, status: ShortListed}
	app2 := &Application{id: 2, status: Applied}
	dr.AppendApplication(app1)
	dr.AppendApplication(app2)
	short := dr.getShortlistedApplications()
	if len(short) != 1 || short[0] != app1 {
		t.Error("getShortlistedApplications failed")
	}
}
