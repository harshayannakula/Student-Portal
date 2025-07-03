package internal

import (
	"fmt"
	"time"
)

var nextID = SeqID()

type JobCategory int

const (
	Day JobCategory = iota
	Dream
	SuperDream
	Marquee
)

var JobCategoryStringMap = map[JobCategory]string{
	Day:        "Day Company",
	Dream:      "Dream",
	SuperDream: "Super Dream",
	Marquee:    "Marquee",
}

func (jc JobCategory) String() string {
	return JobCategoryStringMap[jc]
}

type Eligibility struct {
	requirement float64
}

type Drive struct {
	id           int
	startDate    time.Time
	endDate      time.Time
	roleName     string
	eligibility  Eligibility
	ctc          int
	jobCategory  JobCategory
	applications []*Application
}

func NewEligibility(minimumGPA float64) Eligibility {
	return Eligibility{requirement: minimumGPA}
}

func NewDrive(startDate time.Time, endDate time.Time, roleName string, minimumGPA float64, ctc int, jobCategory JobCategory) Drive {
	return Drive{id: nextID(), startDate: startDate, endDate: endDate, roleName: roleName, eligibility: NewEligibility(minimumGPA), ctc: ctc, jobCategory: jobCategory}
}

// Elegibility setters and Getters
func (el Eligibility) Requirement() float64 {
	return el.requirement
}

func (el Eligibility) ChangeRequirement(newReq float64) {
	el.requirement = newReq
}

// Getters
func (dr Drive) ID() int {
	return dr.id
}

func (dr Drive) StartDate() time.Time {
	return dr.startDate
}

func (dr Drive) EndDate() time.Time {
	return dr.endDate
}

func (dr Drive) RoleName() string {
	return dr.roleName
}

func (dr Drive) Eligibility() Eligibility {
	return dr.eligibility
}

func (dr Drive) CTC() int {
	return dr.ctc
}

func (dr Drive) JobCategory() JobCategory {
	return dr.jobCategory
}

func (dr Drive) Applications() []*Application {
	return dr.applications
}

// Setters
func (dr *Drive) SetStartDate(startDate time.Time) {
	dr.startDate = startDate
}

func (dr *Drive) SetEndDate(endDate time.Time) {
	dr.endDate = endDate
}

func (dr *Drive) SetRoleName(roleName string) {
	dr.roleName = roleName
}

func (dr *Drive) SetEligibility(minimumGPA float64) {
	dr.eligibility.ChangeRequirement(minimumGPA)
}

func (dr *Drive) SetCTC(ctc int) {
	dr.ctc = ctc
}

func (dr *Drive) SetJobCategory(jobCat JobCategory) {
	dr.jobCategory = jobCat
}

func (dr *Drive) AppendApplication(application *Application) {
	dr.applications = append(dr.applications, application)
}

// Drive Functions
func (dr *Drive) HasApplied(StudentID int) bool {
	for _, e := range dr.applications {
		if e.Student.ID() == StudentID {
			return true
		}
	}
	return false
}

func (dr *Drive) GetApplicationByID(id int) (*Application, error) {
	for _, e := range dr.applications {
		if e.ID() == id {
			return e, nil
		}
	}
	return nil, fmt.Errorf("no such application for given id")
}

// Elegibility functions
func (el *Eligibility) checkEligibility(applicant *Applicant) bool {
	if el.requirement > applicant.CGPA {
		return false
	} else {
		return true
	}
}
