package internal

import (
	"encoding/json"
	"testing"
)

// --- Sample Student Constructor for Tests ---
func createSampleStudents() []Student {
	return []Student{
		NewStudent(1, "Alice"),
		NewStudent(2, "Bob"),
		NewStudent(3, "Charlie"),
	}
}

// --- Test NewStudent ---
func TestNewStudent_Valid(t *testing.T) {
	s := NewStudent(10, "Test")
	if s.ID() != 10 || s.Name() != "Test" {
		t.Errorf("expected ID=10, Name='Test'; got ID=%d, Name=%s", s.ID(), s.Name())
	}
}

func TestNewStudent_Invalid(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic for id <= 0, but no panic occurred")
		}
	}()
	NewStudent(0, "Invalid")
}

// --- Test Display ---
func TestDisplay(t *testing.T) {
	s := NewStudent(42, "TestName")
	// not asserting anything here since Display only prints to stdout
	s.Display()
}

// --- Test UpdateStudentName ---
func TestUpdateStudentName(t *testing.T) {
	students := createSampleStudents()
	err := UpdateStudentName(students, 2, "Bobby")
	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
	if students[1].Name() != "Bobby" {
		t.Errorf("expected name to be 'Bobby', got: %s", students[1].Name())
	}

	err = UpdateStudentName(students, 999, "Ghost")
	if err == nil {
		t.Error("expected error for non-existent student ID")
	}
}

// --- Test FindStudentByID ---
func TestFindStudentByID(t *testing.T) {
	students := createSampleStudents()

	s := FindStudentByID(students, 3)
	if s == nil || s.Name() != "Charlie" {
		t.Error("expected to find Charlie")
	}

	s = FindStudentByID(students, 999)
	if s != nil {
		t.Error("expected nil for non-existent student")
	}
}

// --- Test FindStudentsByName ---
func TestFindStudentsByName(t *testing.T) {
	students := createSampleStudents()
	students = append(students, NewStudent(4, "Alice")) // duplicate name

	found := FindStudentsByName(students, "Alice")
	if len(found) != 2 {
		t.Errorf("expected 2 students named Alice, got: %d", len(found))
	}

	found = FindStudentsByName(students, "Unknown")
	if len(found) != 0 {
		t.Errorf("expected 0 results for unknown name, got: %d", len(found))
	}
}

// --- Test DeleteStudentByID ---
func TestDeleteStudentByID(t *testing.T) {
	students := createSampleStudents()

	newList, err := DeleteStudentByID(students, 2)
	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
	if len(newList) != len(students)-1 {
		t.Errorf("expected length %d, got %d", len(students)-1, len(newList))
	}
	for _, s := range newList {
		if s.ID() == 2 {
			t.Error("student with ID 2 was not deleted")
		}
	}

	newList, err = DeleteStudentByID(newList, 999)
	if err == nil {
		t.Error("expected error when deleting non-existent student")
	}
}

// --- Test SerializeStudents ---
func TestSerializeStudents(t *testing.T) {
	students := createSampleStudents()

	data, err := SerializeStudents(students)
	if err != nil {
		t.Errorf("serialization failed: %v", err)
	}
	if len(data) == 0 {
		t.Error("expected non-empty JSON data")
	}

	var back []Student
	err = json.Unmarshal(data, &back)
	if err != nil {
		t.Errorf("unmarshal back failed: %v", err)
	}
}

// --- Test DeserializeStudents ---
func TestDeserializeStudents(t *testing.T) {
	jsonData := `[{"id":1,"name":"Alice"},{"id":2,"name":"Bob"}]`

	students, err := DeserializeStudents([]byte(jsonData))
	if err != nil {
		t.Errorf("deserialization failed: %v", err)
	}
	if len(students) != 2 || students[0].Name() != "Alice" {
		t.Error("unexpected deserialized student data")
	}

	badData := []byte(`{invalid json}`)
	_, err = DeserializeStudents(badData)
	if err == nil {
		t.Error("expected error for invalid JSON")
	}
}
