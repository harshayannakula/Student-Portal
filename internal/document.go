package internal

import "time"

type Document struct {
	Title      string
	Filename   string // e.g. "assignment.pdf"
	Content    []byte // the raw file data
	MimeType   string // e.g. "application/pdf"
	UploadedAt time.Time
}
