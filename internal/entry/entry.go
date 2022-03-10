package entry

import (
	"fmt"
	"time"
)

type Entry struct {
	Description string
	Date        string
	status      bool // false: to do; true: done
}

func (e *Entry) ShowEntry(index int) {
	fmt.Printf("%d(%v): %s \n", index, e.Date, e.Description)
}

func NewEntry(description string) *Entry {
	t := time.Now().Format("2006-01-02")
	return &Entry{description, t, false}
}

func (e *Entry) GetterStatus() bool {
	return e.status
}

func (e *Entry) SetStatus(statusChanged bool) {
	e.status = statusChanged
}
