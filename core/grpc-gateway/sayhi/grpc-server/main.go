package main

import (
	pb "awesomeGo/core/grpc-gateway/sayhi/proto"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	Address = "127.0.0.1:8080"
)

type helloService struct {
}

func (h helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	resp := new(pb.HelloReply)
	resp.Message = "hello," + in.Name + "."
	return resp, nil
}

var HelloServer = helloService{}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		fmt.Printf("failed to listen:%v", err)
	}
	//实现gRPC Server
	s := grpc.NewServer()
	//注册helloServer为客户端提供服务
	pb.RegisterHelloServer(s, HelloServer) //内部调用了s.RegisterSer()
	log.Println("Rpc服务启动成功：Listen on" + Address)
	s.Serve(listen)
}
