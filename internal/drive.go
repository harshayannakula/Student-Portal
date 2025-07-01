package internal

import (
	"fmt"
	"time"
)

type JobCategory int

const (
	Day JobCategory = iota
	Dream
	SuperDream
	Marquee
)

type Drive struct {
	id           int
	startDate    time.Time
	endDate      time.Time
	roleName     string
	eligibility  Eligibility
	ctc          int
	jobCategory  JobCategory // its an enum
	Applications []Application
}

type Eligibility struct {
	requirement int
}

func (pr *PlacementRegistrar) AddDriveToCompany(companyID int, drive Drive) error {
	for i := range pr.companies {
		if pr.companies[i].id == companyID {
			pr.companies[i].drives = append(pr.companies[i].drives, drive)
			return nil

		}
	}
	return fmt.Errorf("company with id %d not found", companyID)
}
