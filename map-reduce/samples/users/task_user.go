package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	coordinator "mit-6.824/map-reduce/proto"
	"time"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:4789", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := coordinator.NewCoordinatorServerClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.QueryTask(ctx, &coordinator.QueryTaskRequest{
		TaskId: "4fc59152-b3f3-f969-b15d-e0e4a3e61bbf",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %v", r)
}
