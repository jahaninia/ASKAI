package main

import (
	"log"
	"github.com/jahaninia/ASKAI/ari"
)

func main() {
	log.Println("Starting ASKAI AI Agent...")

	conn, err := ari.ConnectARI()
	if err != nil {
		log.Fatalf("Failed to connect to ARI: %v", err)
	}
	defer conn.Close()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading ARI message: %v", err)
			continue
		}

		go ari.HandleCallEvent(msg)
	}
}