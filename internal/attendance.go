package internal

import (
	"Student-Portal/Student"
	"time"
)

type AttendanceService struct {
	Records []Student.AttendanceRecord
}

func (a *AttendanceService) MarkAttendance(studentID, sessionID string, present bool) {
	record := Student.AttendanceRecord{
		StudentID: studentID,
		SessionID: sessionID,
		Date:      time.Now(),
		Present:   present,
	}
	a.Records = append(a.Records, record)
}
