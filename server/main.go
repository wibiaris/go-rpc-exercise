package main

import (
	"fmt"
	"log"
	"net"

	"github.com/go-rpc-exercise/api"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}

	s := api.Server{}

	grpcServer := grpc.NewServer()
	api.RegisterPingServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}

}
