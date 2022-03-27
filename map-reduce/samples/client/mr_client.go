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
	r, err := c.CreateTask(ctx, &coordinator.CreateMapReduceTaskRequest{
		Name:       "first map-reduce task",
		Inputs:     []string{"D:\\codes\\go\\src\\github.com\\ctlove0523\\mit-6.824\\map-reduce\\pg-grimm.txt", "D:\\codes\\go\\src\\github.com\\ctlove0523\\mit-6.824\\map-reduce\\pg-huckleberry_finn.txt"},
		MapSize:    2,
		ReduceSize: 3,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetId())
}
