package main

import (
	"context"
	"fmt"
	hello "github.com/leoshus/proto-demo/pb"
	"google.golang.org/grpc"
	"log"
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
	listener, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Printf("start server listen error:%v", err)
		return
	}
	log.Println("start server...")
	if err := server.Serve(listener); err != nil {
		log.Printf("start server error:%v", err)
	}
}
