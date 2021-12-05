package log_process

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) GetLogsByLogLevel(ctx context.Context, req *LogLevelRequest) (*LogLevelReply, error) {
	log.Printf("Receive message body from client: %s %s", req.LogLevel, req.LogPath)
	return &LogLevelReply{Message: "Hello From the Server!"}, nil
}
