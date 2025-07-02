package internal

type Teacher struct {
	ID    string
	Name  string
	Email string
}

// TID returns the Teacher's ID.
// Used for comparisons in attendance logic.
func (t Teacher) TID() string {
	return t.ID
}

// Optional: Constructor for consistency
func NewTeacher(id, name string) Teacher {
	return Teacher{
		ID:   id,
		Name: name,
	}
}
