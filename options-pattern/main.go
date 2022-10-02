package main

import (
	"options-pattern/pkg/server"

	"fmt"
	"time"
)

func main() {
	newServer := server.NewServer("4000", server.WithTimeout(50*time.Second), server.WithMaxConnections(500))

	fmt.Printf("server: %#v", newServer)
}