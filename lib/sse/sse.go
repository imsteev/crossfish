package sse

import "fmt"

type Event struct {
	Type *string
	Data *string
}

func (e Event) Format() string {
	var s string

	if e.Type != nil {
		s += fmt.Sprintf("event:%s\n\n", *e.Type)
	}
	if e.Data != nil {
		s += fmt.Sprintf("data:%s\n\n", *e.Data)
	}
	return s
}