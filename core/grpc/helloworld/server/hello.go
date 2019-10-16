package main

import (
	"awesomeGo/core/grpc/helloworld/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type HelloServiceImpl struct{}

func (h *HelloServiceImpl) Hello(ctx context.Context, req *helloworld.Req) (res *helloworld.Res, err error) {
	fmt.Println(req.GetValue())
	value := "server: Hello World!"
	return &helloworld.Res{Value: value}, nil
}

func main() {
	server := grpc.NewServer()
	helloworld.RegisterHelloServiceServer(server, &HelloServiceImpl{})
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("err:", err)
	}
	server.Serve(lis)
}
