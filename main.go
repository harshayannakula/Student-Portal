package main

import (
	"Student-portal/internal"
	"fmt"
	"time"
)

func main() {
	// Original Registrar
	originalRegistrar := &internal.NewRegistrar{}

	// Wrap in RegistrarWithDocs
	registrar := &internal.RegistrarWithDocs{
		NewRegistrar: originalRegistrar,
	}

	// -------------------------------
	// Create teachers
	// -------------------------------
	teacher1 := internal.NewTeacher("T101", "John Smith")
	teacher2 := internal.NewTeacher("T202", "Jane Doe")

	registrar.AddTeacher(teacher1)
	registrar.AddTeacher(teacher2)

	// -------------------------------
	// Create students
	// -------------------------------
	student1 := internal.NewStudent(1001, "Alice")
	student2 := internal.NewStudent(1002, "Bob")

	registrar.AddStudent(student1)
	registrar.AddStudent(student2)

	// -------------------------------
	// Create courses
	// -------------------------------
	course1 := internal.NewCreditCourse(101, "Data Structures", 4).Course
	course2 := internal.NewCreditCourse(102, "Algorithms", 3).Course

	registrar.AddCourse(course1)
	registrar.AddCourse(course2)

	// -------------------------------
	// Enroll students
	// -------------------------------
	// Alice in Data Structures with teacher1
	enroll1 := internal.NewEnrollnew(
		student1,
		course1,
		internal.PercentageGrader{},
		0.85,
		internal.Attendance{},
		teacher1,
	)
	registrar.Enrollnew(enroll1)

	// Bob in Algorithms with teacher2
	enroll2 := internal.NewEnrollnew(
		student2,
		course2,
		internal.PercentageGrader{},
		0.92,
		internal.Attendance{},
		teacher2,
	)
	registrar.Enrollnew(enroll2)

	// -------------------------------
	// Mark attendance
	// -------------------------------
	// Alice: present on Day 1, absent on Day 2
	day1 := time.Date(2025, 7, 2, 0, 0, 0, 0, time.UTC)
	day2 := day1.AddDate(0, 0, 1)

	internal.Giveattendence(originalRegistrar, 101, 1001, teacher1.TID(), true, day1)
	internal.Giveattendence(originalRegistrar, 101, 1001, teacher1.TID(), false, day2)

	// Bob: present on Day 1, present on Day 2
	internal.Giveattendence(originalRegistrar, 102, 1002, teacher2.TID(), true, day1)
	internal.Giveattendence(originalRegistrar, 102, 1002, teacher2.TID(), true, day2)

	// -------------------------------
	// Display attendance for each student
	// -------------------------------
	fmt.Println("✅ Displaying Attendance Records:\n")

	// Teacher 1 views Alice
	ts1 := &internal.TeacherService{
		Teacher:   teacher1,
		Registrar: registrar,
	}
	fmt.Println("Teacher:", teacher1.Name)
	ts1.DisplayAttendance(101, 1001)
	fmt.Println()

	// Teacher 2 views Bob
	ts2 := &internal.TeacherService{
		Teacher:   teacher2,
		Registrar: registrar,
	}
	fmt.Println("Teacher:", teacher2.Name)
	ts2.DisplayAttendance(102, 1002)
	fmt.Println()
}
