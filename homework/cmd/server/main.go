package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/haichaom/golang-geeke/homework/api/log_process"
)

func main() {

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	server := log_process.Server{}
	log_process.RegisterLogProcessServer(grpcServer, &server)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
