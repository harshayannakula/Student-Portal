package internal

type CourseResult struct {
	StudentId  int             `json:"student_id"`
	CourseId   int             `json:"course_id"`
	CourseName string          `json:"course_name"`
	Grade      AlphabeticGrade `json:"grade"`
	Semester   int             `json:"semester"`
	Credits    float64         `json:"credits"`
}

func NewCourseResult(student_id int, course_id int, courseName string, grade AlphabeticGrade, semester int, credits float64) CourseResult {
	return CourseResult{
		StudentId:  student_id,
		CourseId:   course_id,
		CourseName: courseName,
		Grade:      grade,
		Semester:   semester,
		Credits:    credits,
	}
}
