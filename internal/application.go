package internal

type ApplicationStatus int

const (
	Applied ApplicationStatus = iota

)
type Application struct {
	id int
	Drive
	Applicant
	status ApplicationStatus 
}
