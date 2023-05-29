package main

import (
	"fmt"
	"net/http"
	"time"

	"animalcrossing/fish"
)

// Initialize the SSEMana/ger
var manager = SSEManager{
	clients:    make(map[http.ResponseWriter]chan<- string),
	register:   make(chan clientRegistration),
	unregister: make(chan http.ResponseWriter),
}

func main() {
	http.HandleFunc("/server/events", handleSSE)
	http.ListenAndServe(":8080", nil)
}

func init() {
	go manager.run()
}

func handleSSE(w http.ResponseWriter, r *http.Request) {
	// Set the response headers for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Create a new channel to send messages to the SSE client
	messageChan := make(chan string)

	// Add the client's response writer to the SSEManager
	manager.addClient(w, messageChan)

	// Close the connection when the client closes the connection
	go func() {
		<-r.Context().Done()
		fmt.Println("lost connection")

		manager.removeClient(w)
	}()

	spawner := fish.Spawner{}

	// Send a keep-alive message every 30 seconds to prevent the connection from closing
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		fmt.Println("hello")
		fmt.Fprintf(w, "data: %s\n\n", spawner.Spawn().Name)
		w.(http.Flusher).Flush()
	}
}

// SSEManager is responsible for managing SSE clients and sending messages
type SSEManager struct {
	clients    map[http.ResponseWriter]chan<- string
	register   chan clientRegistration
	unregister chan http.ResponseWriter
}

type clientRegistration struct {
	client   http.ResponseWriter
	messages chan<- string
}

func (m *SSEManager) addClient(client http.ResponseWriter, messages chan<- string) {
	m.register <- clientRegistration{client: client, messages: messages}
}

func (m *SSEManager) removeClient(client http.ResponseWriter) {
	m.unregister <- client
}

func (m *SSEManager) run() {
	for {
		select {
		case registration := <-m.register:
			m.clients[registration.client] = registration.messages
		case client := <-m.unregister:
			delete(m.clients, client)
		}
	}
}
