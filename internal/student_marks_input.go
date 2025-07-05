package internal

type StudentMarkInput struct {
	CourseID  int     `json:"course_id"`
	StudentID int     `json:"student_id"`
	Score     float64 `json:"score"`
}
