package internal

import "fmt"

type PlacementRegistrar struct {
	companies    []Company
	applications []Application
	applicants   []Applicant
}

type ReportByStudent struct {
	Applicant
	totalOffersRecived []Application
	eligibileRoles     []Drive
	finalOffer         Applicant
	ctcForFinalOffer   int
}

type FullPlacementReport struct {
	totalComapanies       []Company
	totalOffersMade       int
	totalOffersByCatagory map[JobCategory]int
}

func (pr PlacementRegistrar) GenerateReportByStudent() ReportByStudent {
	return ReportByStudent{}
}

func (pr PlacementRegistrar) GenerateFullReport() FullPlacementReport {
	return FullPlacementReport{}
}

func (pr *PlacementRegistrar) AddCompany(company Company) { //AddCompany will help us add new company to PlacementRegistrar
	pr.companies = append(pr.companies, company)
}

func (pr *PlacementRegistrar) GetCompanyByID(id int) (*Company, error) {
	for i := range pr.companies {
		if pr.companies[i].id == id {
			return &pr.companies[i], nil
		}
	}
	return nil, fmt.Errorf("company with id %d is not found", id)
}

func (pr *PlacementRegistrar) UpdateCompany(UpdatedCompany Company) error {
	for i := range pr.companies {
		if pr.companies[i].id == UpdatedCompany.id {
			pr.companies[i] = UpdatedCompany
			return nil
		}
	}
	return fmt.Errorf("Company not found %d to update", UpdatedCompany.id)
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

func (pr *PlacementRegistrar) ApplicantByID(studentID int) (*Applicant, error) {
	for i := range pr.applicants {
		if pr.applicants[i].ID() == studentID {
			return &pr.applicants[i], nil
		}
	}
	return nil, fmt.Errorf("applicant with id %d not found", studentID)
}
