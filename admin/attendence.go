package admin

import (
	"time"
)

type Attendance struct {
	Records map[time.Time]bool
}
