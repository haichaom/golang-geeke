package main

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/haichaom/golang-geeke/homework/api/log_process"
	"github.com/haichaom/golang-geeke/homework/errors"
)

type LogServer struct {
	pb.UnimplementedLogProcessServer
}

func (s *LogServer) GetLogsByLogLevel(ctx context.Context, req *pb.LogLevelRequest) (*pb.LogLevelReply, error) {
	log.Printf("Receive message body from client: %s %s", req.LogLevel, req.LogPath)
	if req.LogPath == "invalid.log" {
		return nil, errors.BadRequest("invalid request",
			fmt.Sprintf("invalid argument logpath: %s", req.LogPath))
	}
	return &pb.LogLevelReply{Message: "Hello From the Server!"}, nil
}

func main() {

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	server := LogServer{}
	pb.RegisterLogProcessServer(grpcServer, &server)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
