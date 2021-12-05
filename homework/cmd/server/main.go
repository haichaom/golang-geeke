package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"golang.org/x/net/context"

	pb "github.com/haichaom/golang-geeke/homework/api/log_process"
)

type LogServer struct {
    pb.UnimplementedLogProcessServer
}

func (s *LogServer) GetLogsByLogLevel(ctx context.Context, req *pb.LogLevelRequest) (*pb.LogLevelReply, error) {
	log.Printf("Receive message body from client: %s %s", req.LogLevel, req.LogPath)
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
