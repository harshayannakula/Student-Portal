package internal

import "fmt"

type Company struct {
	id     int
	name   string
	drives []Drive
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
	return nil, fmt.Errorf("company with id %d is not found.", id)
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
