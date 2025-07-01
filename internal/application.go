package internal

type ApplicationStatus int

const (
	Applied ApplicationStatus = iota
	ShortListed
	Cleared
	Selected
	Rejected
)

type Application struct {
	id int
	Drive
	Applicant
	status ApplicationStatus
}
