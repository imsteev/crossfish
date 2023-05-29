package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"crossfish/fish"
	"crossfish/lib/sse"
)

const SPAWN_EVERY_SECONDS = 1

func main() {
	http.HandleFunc("/server/events", handleSSE)
	http.ListenAndServe(":8080", nil)
}

func handleSSE(w http.ResponseWriter, r *http.Request) {
	// Set the response headers for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	ticker := time.NewTicker(SPAWN_EVERY_SECONDS * time.Second)
	defer ticker.Stop()

	// cleanup resources once connection is terminated
	go func() {
		defer ticker.Stop()
		<-r.Context().Done()
		log.Printf("connection terminated: %v\n", r.Context().Err())
	}()

	for range ticker.C {
		fish := fish.SpawnFish()
		typ := new(string)
		*typ = "new-fish"
		e := sse.Event{Data: &fish.Name, Type: typ}
		serialized := e.Format()
		log.Println(serialized)
		fmt.Fprint(w, serialized)
		w.(http.Flusher).Flush()
	}
}
