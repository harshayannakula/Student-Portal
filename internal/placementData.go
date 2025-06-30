package internal

import (
	"time"
)

type companies struct {
	id    int
	name  string
	drive []Drive
	// selectedApplicants

}

type Drive struct {
	id                 int
	startDate          time.Time
	endDate            time.Time
	roleName           string
	eligibility        Eligibility
	ctc                int
	jobCategory        JobCategory // its an enum
	selectedApplicants []Applicant
}

type Eligibility struct {
	requirement int
}

type JobCategory int

const (
	Day JobCategory = iota
	Dream
	SuperDream
	Marquee
)

type Applicant struct {
	Student
	drivesAppliedFor []Drive
}
