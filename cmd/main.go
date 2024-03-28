package main

import (
	"log"
	"net/http"
	"time"

	"github.com/imsteev/crossfish/fish"
	"github.com/imsteev/crossfish/lib/sse"
	"github.com/imsteev/crossfish/middleware"
)

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
	spn = fish.SpawnEvery(time.Second)
	for f := range spn.C {
		srv.WriteEvent("new-fish", f.Name)
		srv.Flush()
	}
}
