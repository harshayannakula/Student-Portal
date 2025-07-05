package internal

type StudentResult struct {
    CourseID   int     `json:"course_id"`
    CourseName string  `json:"course_name"`
    StudentID  int     `json:"student_id"`
    StudentName string `json:"student_name"`
    Score      float64 `json:"score"`
    Grade      string  `json:"grade"`
}
