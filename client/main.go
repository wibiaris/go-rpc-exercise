package main

import (
	"context"
	"log"

	"github.com/go-rpc-exercise/api"

	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":7777", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect : %s", err)
	}

	defer conn.Close()

	c := api.NewPingClient(conn)
	response, err := c.SayHello(context.Background(), &api.PingMessage{Greeting: "foo"})

	if err != nil {
		log.Fatalf("Errror When Calling sayhello : %s", err)
	}
	log.Printf("Response from server %s", response.Greeting)
}
