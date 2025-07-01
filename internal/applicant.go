package internal

type Applicant struct {
	Student
	AcademicRecord
	drivesAppliedFor []Drive
}

func NewApplicant(st Student, ar AcademicRecord) Applicant {
	return Applicant{Student: st, AcademicRecord: ar}
}

func (a *Applicant) DrivesAppliedFor() []Drive {
	return a.drivesAppliedFor
}

func (a *Applicant) AddDrivesAppliedFor(drive Drive) {
	a.drivesAppliedFor = append(a.drivesAppliedFor, drive)
}

func (a *Applicant) SetDrivesAppliedFor(drives []Drive) {
	a.drivesAppliedFor = drives
}

