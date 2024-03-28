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
	var spn *fish.Spawner

	go func() {
		<-r.Context().Done()
		spn.Stop()
		log.Printf("connection terminated: %v\n", r.Context().Err())
	}()

	srv := sse.NewServer(w)
	spn = fish.SpawnEvery(SPAWN_EVERY_SECONDS * time.Second)
	for f := range spn.C {
		srv.WriteEvent("new-fish", f.Name)
		srv.Flush()
	}
}
