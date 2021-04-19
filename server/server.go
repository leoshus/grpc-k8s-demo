package main

import (
	"context"
	"fmt"
	hello "github.com/leoshus/proto-demo/pb"
	"google.golang.org/grpc"
	"net"
	"os"
	"time"
)

type HelloServer struct {
}

func (h *HelloServer) SayHello(ctx context.Context, req *hello.HelloRequest) (*hello.HelloReply, error) {
	now := time.Now().Format("2006-01-02 15:04:05")
	hostname, _ := os.Hostname()
	fmt.Printf("%s say hello:%s\n", hostname, now)
	return &hello.HelloReply{
		Message: fmt.Sprintf("%s say hello %s :%s", hostname, req.Name, now),
	}, nil
}

func main() {
	server := grpc.NewServer()
	hello.RegisterGreeterServer(server, &HelloServer{})
	listener, _ := net.Listen("tcp", "127.0.0.1:8088")
	server.Serve(listener)
}
