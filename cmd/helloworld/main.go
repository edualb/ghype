package main

import (
	"log"

	"github.com/edualb/ghype/internal/conn"
)

func main() {
	// When we type curl localhost:8080,
	// We are receveing the following error: curl: (1) Received HTTP/0.9 when not allowed
	s := conn.NewServer("tcp", ":8080")
	log.Fatal(s.Listen())
}
