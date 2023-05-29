package sse

import "fmt"

type Event struct {
	Type string
	Data string
}

func (e Event) Format() string {
	var s string
	if e.Type != "" {
		s += fmt.Sprintf("event: %s\n", e.Type)
	}
	if e.Data != "" {
		s += fmt.Sprintf("data:%s\n", e.Data)
	}
	s += "\n"
	return s
}
