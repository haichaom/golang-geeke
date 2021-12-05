package main

import (
	"context"
	"log"

	"github.com/haichaom/golang-geeke/homework/api/log_process"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Failed to connect: %s", err)
	}
	defer conn.Close()

	client := log_process.NewLogProcessClient(conn)
	req := &log_process.LogLevelRequest{LogLevel: "Error", LogPath: "/nfs/log/123456/api.log"}
	resp, err := client.GetLogsByLogLevel(context.Background(), req)
	if err != nil {
		log.Fatal("Failed to get logs by loglevel: %s", err)
	}
	log.Printf("Received from server:%s ", resp.Message)
}
