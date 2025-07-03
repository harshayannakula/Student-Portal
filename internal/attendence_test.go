package internal

// import (
// 	"fmt"
// 	"testing"
// 	"time"
// )

// func TestMarkAttendance(t *testing.T) {
// 	attendTime := time.Now()

// 	ok1 := Giveattendence(&Registrar1, Course1.Id, Student1.id, Teacher1.id, true, attendTime)
// 	if !ok1 {
// 		t.Error("Error in Giving attendence")
// 	}

// 	a, ok := FetchAttendance(&Registrar1, Course1.Id, Student1.id, Teacher1.id)
// 	if !ok {
// 		t.Error("Attendance not marked")
// 		fmt.Println(Course1.Id)
// 	}
// 	if a[attendTime] != true {
// 		t.Error("Wrong attendance value")
// 	}
// }
