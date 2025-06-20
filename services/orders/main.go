package main

import (
	"log"

	orderhttp "kitchen.local/services/orders/service"
)

func main() {
	// Start the HTTP server in a goroutine
	httpServer := orderhttp.NewHTTPServer(":8000")
	go func() {
		if err := httpServer.Run(); err != nil {
			log.Fatalf("HTTP server failed: %v", err)
		}
	}()

	// Start the gRPC server
	grpcServer := NewGRPCServer(":9000")
	if err := grpcServer.Run(); err != nil {
		log.Fatalf("gRPC server failed: %v", err)
	}
}
