package main

import (
	"awesomeGo/core/grpc/go-grpc-example/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type SearchServer struct {
}

func (s *SearchServer) Search(ct context.Context, req *proto.SearchReq) (*proto.SearchRes, error) {
	fmt.Println("service search function")
	res := &proto.SearchRes{Res: req.GetReq() + "Server"}
	return res, nil
}

func main() {
	searchServer := SearchServer{}
	server := grpc.NewServer()
	proto.RegisterSearchServiceServer(server, &searchServer)
	lis, err := net.Listen("tcp", ":8099")
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	server.Serve(lis)
}
