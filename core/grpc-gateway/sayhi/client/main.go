package main

import (
	pb "awesomeGo/core/grpc-gateway/sayhi/proto"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	Address = "127.0.0.1:8080"
)

func main() {
	//连接gRPC服务器
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	//初始化客户端
	c := pb.NewHelloClient(conn)

	//调用方法
	reqBody := new(pb.HelloRequest)
	reqBody.Name = " grpc."
	r, err := c.SayHello(context.Background(), reqBody)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r.Message)

}
