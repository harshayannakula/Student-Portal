package domain

type ResultRegister struct {
	results []SemesterResult
}

func (r *ResultRegister) SetResults(newResults []SemesterResult) {
	r.results = newResults
}

func (r ResultRegister) Results() []SemesterResult {
	return r.results

}
