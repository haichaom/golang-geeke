package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"

	"golang.org/x/sync/errgroup"
)

func helloWorld(w http.ResponseWriter, req *http.Request) {
	fmt.Println("func-helloWorld:: called")
	io.WriteString(w, "Hello world")
}

func runServer(srv *http.Server) error {
	http.HandleFunc("/hello", helloWorld)
	return srv.ListenAndServe()
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	group, errCtx := errgroup.WithContext(ctx)
	srv := &http.Server{Addr: ":80"}
	group.Go(func() error {
		fmt.Println("main:: start http Server", srv)
		return runServer(srv)
	})
	group.Go(func() error {
		<-errCtx.Done()
		fmt.Println("main:: Shut down httpServer")
		return srv.Shutdown(errCtx)
	})
	c := make(chan os.Signal, 1)
	signal.Notify(c)
	group.Go(func() error {
		for {
			select {
			case <-errCtx.Done():
				fmt.Println("main-select:: ctx Done")
				return errCtx.Err()
			case sig := <-c:
				fmt.Println("main-select:: Signal received: ", sig)
				cancel()
			}
		}
	})

	if err := group.Wait(); err != nil {
		fmt.Println("main:: Found error: ", err)
	} else {
		fmt.Println("main:: Done")
	}
}
