package main

import (
	"context"
	"flag"
	"fmt"
	hello "github.com/leoshus/proto-demo/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/balancer/roundrobin"
	"log"
	"strings"
	"time"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate)
	var address string
	flag.StringVar(&address, "address", "localhost:8088", "grpc server address")
	flag.Parse()
	log.Println("grpc client start begin...")
	conn, err := grpc.Dial(strings.Join([]string{"dns:///", address}, ""), grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy":"%s"}`, roundrobin.Name)),
		//grpc.WithBlock(),
		grpc.WithConnectParams(grpc.ConnectParams{
			Backoff: backoff.Config{
				MaxDelay: 2 * time.Second,
			},
			MinConnectTimeout: 2 * time.Second,
		}))
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Println("grpc client start success")
	defer conn.Close()
	client := hello.NewGreeterClient(conn)
	for range time.Tick(time.Second) {
		resp, err := client.SayHello(context.TODO(), &hello.HelloRequest{
			Name: "tom",
		})
		if err != nil {
			fmt.Println(err)
			log.Printf("say hello occur error:%v\n", err)
			return
		}
		fmt.Println(resp)
		log.Printf("say hello : %s \n", resp)
	}

}
