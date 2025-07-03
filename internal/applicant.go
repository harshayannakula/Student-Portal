package internal

type Applicant struct {
	Student
	AcademicRecord
	drivesAppliedFor []*Drive
}

func NewApplicant(st Student, ar AcademicRecord) Applicant {
	return Applicant{Student: st, AcademicRecord: ar}
}

func (a *Applicant) DrivesAppliedFor() []*Drive {
	return a.drivesAppliedFor
}

func (a *Applicant) AddDrivesAppliedFor(drive *Drive) {
	a.drivesAppliedFor = append(a.drivesAppliedFor, drive)
}

func (a *Applicant) SetDrivesAppliedFor(drives []*Drive) {
	a.drivesAppliedFor = drives
}

func (a *Applicant) CompaniesAppliedFor(pr *PlacementRegistrar) []string {
	companySet := make(map[string]struct{})

	for _, app := range pr.applications {
		if app.Applicant.Student.id == a.Student.id {
			for _, company := range pr.companies {
				for _, drive := range company.drives {
					if drive.id == app.driveId {
						companySet[company.name] = struct{}{}
					}
				}
			}
		}
	}
	companies := make([]string, 0, len(companySet))
	for name := range companySet {
		companies = append(companies, name)
	}
	return companies
}

func (a *Applicant) TotalNumberOfCompaniesAppliedFor(pr *PlacementRegistrar) int {
	return len(a.CompaniesAppliedFor(pr))
}
