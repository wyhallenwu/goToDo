package entry

import (
	"fmt"
	"time"
)

type Entry struct {
	Description string
	Date        string
	status      bool   // false: to do; true: done
	Project     string // an entry belongs to one project
}

func (e *Entry) PrintEntry(index int) {
	fmt.Printf("%d(%v): %s @(%s) \n", index, e.Date, e.Description, e.Project)
}

func NewEntry(description string) *Entry {
	t := time.Now().Format("2006-01-02")
	return &Entry{description, t, false, ""}
}

func (e *Entry) GetterStatus() bool {
	return e.status
}

func (e *Entry) SetStatus(statusChanged bool) {
	e.status = statusChanged
}

func (e *Entry) SetProject(project string) {
	e.Project = project
}
