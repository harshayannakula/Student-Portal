
package internal

import (
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

	id                 int
	startDate          time.Time
	endDate            time.Time
	roleName           string
	eligibility        Eligibility
	ctc                int
	jobCategory        JobCategory // its an enum
	Applications 	   []Application
}

type Eligibility struct {
	requirement int
}
