package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/f4ah6o/direct-go-sdk/direct-go/logserver"
)

func main() {
	port := flag.String("port", "9999", "server port")
	flag.Parse()

	srv := logserver.New()
	addr := ":" + *port

	fmt.Printf("Log server running at http://localhost%s\n", addr)
	fmt.Printf("API endpoints:\n")
	fmt.Printf("  GET /logs    - latest logs (JSON)\n")
	fmt.Printf("  GET /stream  - real-time stream (SSE)\n")

	if err := srv.ListenAndServe(addr); err != nil {
		log.Fatal(err)
	}
}
