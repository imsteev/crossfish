package sse

import (
	"fmt"
	"log"
	"net/http"
)

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

// EventServer writes and sends data over HTTP
type EventServer struct {
	Writer http.ResponseWriter
}

func NewServer(w http.ResponseWriter) *EventServer {
	return &EventServer{w}
}

// WriteEvent is responsible for writing valid protocol to the sink.
func (e EventServer) WriteEvent(typ, data string) {
	serialized := Event{Type: typ, Data: data}.Format()
	log.Println(serialized)
	fmt.Fprint(e.Writer, serialized)
}

func (e EventServer) Flush() {
	e.Writer.(http.Flusher).Flush()
}
