package internal

import (
	"time"
)

type Attendance struct {
	Records map[time.Time]bool
}

func markAttendance(attendance *Attendance, date time.Time, present bool) {
	if attendance.Records == nil {
		attendance.Records = make(map[time.Time]bool)
	}
	attendance.Records[date] = present
}

// need to check time data type parameter passing
func Giveattendence(r *NewRegistrar, courseID int, studentID int, TeacherID string, attendence bool, time time.Time) {
	for _, e := range r.enroll {
		if e.Course.Id == courseID && e.Student.ID() == studentID && e.Teacher.TID() == TeacherID {
			//r.enroll[i].Attendence = attendence
			markAttendance(&e.attend, time, attendence)
			return
		}
	}
}
