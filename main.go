package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/takumi616/go-english-vocabulary-api/config"
)

func run(ctx context.Context) error {
	//Get config (environment variables)
	config, err := config.GetConfig()
	if err != nil {
		log.Printf("Failed to get config: %v", err)
		return err
	}

	//Set up http listener with port from config file
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		log.Fatalf("Failed to get a listener with port %d: %v", config.Port, err)
	}

	//Set up routing
	mux := SetUpRouting()

	//Run server
	server := NewServer(mux, listener)
	return server.Start(ctx)
}

func main() {
	err := run(context.Background())
	if err != nil {
		log.Printf("Main process error: %v", err)
	}
}
