package domain

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type StudentRegister struct {
	Students []Student
}

type StudentData struct{
	ID int `json:"id"`
	Name string `json:"name"`
}

func (s *StudentRegister) LoadStudents() {
    var students []StudentData
    data, err := os.ReadFile("students.json")
    if err != nil {
        log.Fatal("Failed to read students.json:", err)
    }
    
    err = json.Unmarshal(data, &students)
    if err != nil {
        log.Fatal("Failed to unmarshal students:", err)
    }
    s.Students = make([]Student, 0, len(students))
    for _, sd := range students {
	    student := NewStudent(sd.ID, sd.Name)
	    s.Students = append(s.Students, student)
    }
}

func (s StudentRegister) Display() {
	for _, sd := range s.Students {
		fmt.Printf("#%d : %s\n", sd.ID(), sd.Name())
	}
}
