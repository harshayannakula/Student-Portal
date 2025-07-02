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

type LetterGrader struct{}

func (LetterGrader) Grade(e Enrollment) (string, error) {
	switch {
	case e.score >= 9.0:
		return "A", nil
	case e.score >= 8.0:
		return "B", nil
	case e.score >= 7.0:
		return "C", nil
	case e.score >= 6.0:
		return "D", nil
	default:
		return "F", nil
	}
}
