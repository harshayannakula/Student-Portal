package internal

type PlacementRegistrar struct {
	companies []Company
	applications []Application
}

type ReportByStudent struct {
	Applicant
	totalOffersRecived []Application
	eligibileRoles []Drive
	finalOffer Applicant
	ctcForFinalOffer int
}

type FullPlacementReport struct {
	totalComapanies []Company
	totalOffersMade int
	totalOffersByCatagory map[JobCategory]int
}

func (pr PlacementRegistrar) GenerateReportByStudent() ReportByStudent {
	return ReportByStudent{}
}

func (pr PlacementRegistrar) GenerateFullReport() FullPlacementReport {
	return FullPlacementReport{}
}
