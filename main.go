package main

import "log"

// run proxy server, listen for requests from server1 and forward them to server2
func main() {
	if err := RunServer(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
