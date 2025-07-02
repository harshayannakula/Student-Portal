func (ts *TeacherService) UploadDocument(
	courseID int,
	studentID int,
	title string,
	content string,
) error {
	found := false

	for i, e := range ts.Registrar.enroll {
		if e.Course.Id == courseID &&
			e.Student.ID() == studentID &&
			e.Teacher.TID() == ts.Teacher.TID() {

			// Create document
			doc := Document{
				Title:      title,
				Content:    content,
				UploadedAt: time.Now().Format("2006-01-02 15:04:05"),
			}

			// Wrap in EnrollnewWithDocs
			enrollWithDocs := EnrollnewWithDocs{
				Enrollnew: e,
				Documents: []Document{doc},
			}

			fmt.Printf("✅ Document uploaded: '%s' for student %d in course %d by teacher %s\n",
				doc.Title, studentID, courseID, ts.Teacher.TID())

			fmt.Println("📎 Uploaded documents:")
			for _, d := range enrollWithDocs.Documents {
				fmt.Printf("- %s (at %s)\n", d.Title, d.UploadedAt)
			}

			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("no valid enrollment found for this teacher, student, and course")
	}

	return nil
}
