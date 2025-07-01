package internal

import (
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
	requirement int
}

type Drive struct {
	id           int
	startDate    time.Time
	endDate      time.Time
	roleName     string
	eligibility  Eligibility
	ctc          int
	jobCategory  JobCategory
	applications []Application
}

func NewElegibility(minimumGPA int) Eligibility {
	return Eligibility{requirement: minimumGPA}
}

func NewDrive(startDate time.Time, endDate time.Time, roleName string, minimumGPA int, ctc int, jobCategory JobCategory) Drive {
	return Drive{id: nextID(), startDate: startDate, endDate: endDate, roleName: roleName, eligibility: NewElegibility(minimumGPA), ctc: ctc, jobCategory: jobCategory}
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

func (dr Drive) Applications() []Application {
	return dr.applications
}

// Setters
func (dr Drive) SetStartDate() time.Time {
	return dr.startDate
}

func (dr Drive) SetEndDate() time.Time {
	return dr.endDate
}

func (dr Drive) SetRoleName() string {
	return dr.roleName
}

func (dr Drive) SetElegibility() Eligibility {
	return dr.eligibility
}

func (dr Drive) SetCTC() int {
	return dr.ctc
}

func (dr Drive) SetJobCategory() JobCategory {
	return dr.jobCategory
}

func (dr *Drive) AppendApplication(application Application) {
	dr.applications = append(dr.applications, application)
}
