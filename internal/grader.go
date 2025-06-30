package internal

import "fmt"

type Grader interface {
	Grade(e Enrollment) (string, error)
}

type PercentageGrader struct{}


func (PercentageGrader) Grade(e Enrollment) (string, error) {
	return fmt.Sprintf("%.1f%%", e.score*100), nil
}

type PassFailGrader struct {
	PassMark float64
}

func (p PassFailGrader) Grade(e Enrollment) (string, error) {
	if e.score >= p.PassMark {
		return "PASS", nil
	} else {
		return "FAIL", nil
	}
}
