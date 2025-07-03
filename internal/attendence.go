package internal

import (
	"time"
)

type Attendance struct {
	Records map[time.Time]bool
}

// function to give attendence
func markAttendance(attendance *Attendance, date time.Time, present bool) {
	if attendance.Records == nil {
		attendance.Records = make(map[time.Time]bool)
	}
	attendance.Records[date] = present
}

// function to give attendance to a student in a course by a teacher
// This function will be called by the teacher to mark attendance for a student in a course.
// need to check time data type parameter passing
func Giveattendence(r *NewRegistrar, courseID int, studentID int, TeacherID int, attendence bool, time time.Time) {
	for _, e := range r.enroll {
		if e.Course.Id == courseID && e.Student.ID() == studentID && e.Teacher.ID() == TeacherID {
			//r.enroll[i].Attendence = attendence
			markAttendance(&e.attend, time, attendence)
			return
		}
	}
}

//fetching useful for both student and teacher

func fetchAttendance(r *NewRegistrar, courseID int, studentID int, TeacherID int) (map[time.Time]bool, bool) {
	for _, e := range r.enroll {
		if e.Course.Id == courseID && e.Student.ID() == studentID && e.Teacher.ID() == TeacherID {
			return e.attend.Records, true
		}
	}
	return nil, false
}
