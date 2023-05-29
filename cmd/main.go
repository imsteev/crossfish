package main

import (
	"log"
	"net/http"
	"time"

	"crossfish/fish"
	"crossfish/lib/sse"
	"crossfish/middleware"
)

const SPAWN_EVERY_SECONDS = 1

func main() {
	http.HandleFunc("/server/events", middleware.ServerSentEvents(handleSSE))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func handleSSE(w http.ResponseWriter, r *http.Request) {
	ticker := time.NewTicker(SPAWN_EVERY_SECONDS * time.Second)
	defer ticker.Stop()

	// cleanup resources once connection is terminated
	go func() {
		defer ticker.Stop()
		<-r.Context().Done()
		log.Printf("connection terminated: %v\n", r.Context().Err())
	}()

	eventServer := sse.EventServer{Writer: w}
	for range ticker.C {
		fish := fish.Spawn()
		eventServer.Write(sse.Event{Type: "new-fish", Data: fish.Name})
		eventServer.Flush()
	}
}
