package internal

import (
	"encoding/json"
	"errors"
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

func (ts *TeacherService) UploadStudentMark(courseID int, studentID int, score float64) error {
	for i, e := range ts.Registrar.enroll {
		if e.Course.Id == courseID && e.Student.ID() == studentID && e.Teacher.TID() == ts.Teacher.TID() {
			ts.Registrar.enroll[i].score = score
			grade, err := e.Grader.Grade(ts.Registrar.enroll[i].Enrollment)
			if err != nil {
				return err
			}
			fmt.Printf("Uploaded score %.2f for student %d in course %d. Grade: %s\n", score, studentID, courseID, grade)
			return nil
		}
	}
	return errors.New("no valid enrollment found for this teacher, student, and course")
}

func (ts *TeacherService) UploadStudentMarksFromJSON(jsonData []byte) error {
	var marks []StudentMarkInput
	if err := json.Unmarshal(jsonData, &marks); err != nil {
		return err
	}

	var errs []string
	for _, mark := range marks {
		err := ts.UploadStudentMark(mark.CourseID, mark.StudentID, mark.Score)
		if err != nil {
			errs = append(errs, fmt.Sprintf("student %d, course %d: %v", mark.StudentID, mark.CourseID, err))
		}
	}

	if len(errs) > 0 {
		return errors.New("Errors occurred: " + fmt.Sprint(errs))
	}
	return nil
}

func (ts *TeacherService) GetCourseResults(courseID int) ([]StudentResult, error) {
    var results []StudentResult
    for _, e := range ts.Registrar.enroll {
        if e.Course.Id == courseID && e.Teacher.TID() == ts.Teacher.TID() {
            grade, _ := e.Grader.Grade(e.Enrollment)
            results = append(results, StudentResult{
                CourseID:    e.Course.Id,
                CourseName:  e.Course.Name,            
                StudentID:   e.Student.ID(),
                StudentName: e.Student.Name(),      
                Score:       e.score,
                Grade:       grade,
            })
        }
    }
    if len(results) == 0 {
        return nil, fmt.Errorf("no students found for course %d and teacher %d\n", courseID, ts.Teacher.TID())
    }
    return results, nil
}

