package main

import (
	"awesomeGo/core/grpc/go-grpc-example/proto"
	"context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial(":8099", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()

	client := proto.NewSearchServiceClient(conn)

	req := &proto.SearchReq{
		Req: "grpc client ",
	}

	res, err := client.Search(context.Background(), req)
	if err != nil {
		log.Fatalf("client.Search err: %v", err)
	}
	log.Printf("resp: %s", res.GetRes())
}
