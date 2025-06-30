package internal

type StudentGPA struct {
	Student  Student
	Semester int
	Gpa      float64 // GPA out of 10
}

type StudentRecord struct {
	Student    Student
	Semesters  []StudentGPA
	OverallGPA float64
	Status     string // "At Risk", "Dean's List", "Normal"
}

type GPACalculator struct{}

func NewGPACalculator() *GPACalculator {
	return &GPACalculator{}
}

func (gc *GPACalculator) CalculateOverallGPA(semesters []StudentGPA) float64 {
	if len(semesters) == 0 {
		return 0.0
	}

	total := 0.0
	for _, semester := range semesters {
		total += semester.Gpa
	}
	return total / float64(len(semesters))
}

func (gc *GPACalculator) DetermineStatus(overallGPA float64) string {
	if overallGPA <= 2.0 {
		return "At Risk"
	} else if overallGPA > 8.0 {
		return "Dean's List"
	}
	return "Normal"
}
