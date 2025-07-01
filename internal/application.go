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

func (app *Application) ID() int {
	return app.id
}

func (app *Application) Status() ApplicationStatus {
	return app.status
}
