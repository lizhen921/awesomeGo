package main

import (
	"awesomeGo/core/grpc/helloworld/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := helloworld.NewHelloServiceClient(conn)
	res, err := client.Hello(context.Background(), &helloworld.Req{Value: "client: helloworld"})
	if err != nil {
		fmt.Printf("请求err： %v", err)
	}
	fmt.Println(res)
}
