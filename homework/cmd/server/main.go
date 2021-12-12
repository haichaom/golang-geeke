package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	pb "github.com/haichaom/golang-geeke/homework/api/log_process"
	"github.com/haichaom/golang-geeke/homework/errors"
	"github.com/haichaom/golang-geeke/homework/pkg/cache"
)

const (
	DEFAULT_LOG_MSG = "Hello From the Server!"
)

var (
	grpcServer *grpc.Server
)
var logCache = cache.New("logCache")

type LogServer struct {
	pb.UnimplementedLogProcessServer
}

func (s *LogServer) GetLogsByLogLevel(ctx context.Context, req *pb.LogLevelRequest) (*pb.LogLevelReply, error) {
	log.Printf("Receive message body from client: %s %s", req.LogLevel, req.LogPath)
	if req.LogPath == "invalid.log" {
		return nil, errors.BadRequest("invalid_request",
			fmt.Sprintf("invalid argument logpath: %s", req.LogPath))
	}
	message, isOk := logCache.Get(req.LogPath)
	if !isOk {
                // ..omit log process logic
		message = DEFAULT_LOG_MSG
		logCache.Set(req.LogPath, message)
	}
	return &pb.LogLevelReply{Message: message}, nil
}

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	group, errCtx := errgroup.WithContext(ctx)
	group.Go(func() error {
		lis, err := net.Listen("tcp", ":9000")
		if err != nil {
			return errors.InternalServer("failed_start_server",
				fmt.Sprintf("failed to listen : %s", err))
		}

		grpcServer = grpc.NewServer()
		server := LogServer{}
		pb.RegisterLogProcessServer(grpcServer, &server)

		if err := grpcServer.Serve(lis); err != nil {
			return errors.InternalServer("failed_start_server",
				fmt.Sprintf("failed to serve: %s", err))
		}
		return nil
	})
	group.Go(func() error {
		<-errCtx.Done()
		if grpcServer != nil {
			log.Print("Shutdown grpc server")
			grpcServer.GracefulStop()
		}
		return nil
	})
	group.Go(func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c)
		for {
			select {
			case <-errCtx.Done():
				log.Print("ctx is Done in listening interuppt routine")
				return errCtx.Err()
			case sig := <-c:
				log.Printf("Signal received: ", sig)
				cancel()
			}
		}
	})

	if err := group.Wait(); err != nil {
		log.Printf("grpcServer run into an error: ", err)
	} else {
		log.Print("All is done gracfully")
	}
}
