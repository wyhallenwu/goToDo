package entry

import (
	"fmt"
	"time"
)

type Entry struct {
	index       int
	description string
	date        time.Time
}

func (e *Entry) ShowEntry() {
	fmt.Printf("%d %s %v", e.index, e.description, e.date)
}

func NewEntry(index int, description string, date time.Time) *Entry {
	return &Entry{index, description, date}
}
