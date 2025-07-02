package internal

import (
	"fmt"
)

type TeacherService struct {
	Registrar *RegistrarWithDocs
	Teacher   Teacher
}

func (ts *TeacherService) DisplayAttendance(courseID int, studentID int) {
	for _, e := range ts.Registrar.enroll {
		if e.Course.Id == courseID &&
			e.Student.ID() == studentID &&
			e.Teacher.TID() == ts.Teacher.TID() {
			fmt.Printf("Attendance for student %d in course %d:\n", studentID, courseID)
			for date, present := range e.Attend.Records {
				fmt.Println(date.Format("2006-01-02"), "->", present)
			}
			return
		}
	}
	fmt.Println("No attendance records found.")
}
