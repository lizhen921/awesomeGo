package main

import (
	pb "awesomeGo/core/grpc/go-grpc-example/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {
	//连接grpc server
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("err：", err)
	}
	defer conn.Close()
	client := pb.NewStreamServiceClient(conn)
	//服务端流式RPC
	log.Println("Demo1: 服务端流式RPC")
	ServerSideStreamDemo(client)
	time.Sleep(time.Second)
	//客户端流式RPC
	log.Println("Demo2: 客户端流式RPC")
	ClientSideStreamDemo(client)
	time.Sleep(time.Second)
	//双向流式RPC
	log.Println("Demo3: 双向流式RPC")
	DoubleSideStreamDemo(client)
}

//客户端发起一次请求，服务端流式返回
func ServerSideStreamDemo(client pb.StreamServiceClient) error {
	req := &pb.StreamReq{
		Pt: &pb.StreamPoint{
			Name:  "Normal Client 请求 Stream Server",
			Value: 1,
		},
	}
	fmt.Printf("Normal Client 请求 Stream Server\n")
	stream, err := client.ServerSideStream(context.Background(), req)
	if err != nil {
		fmt.Println("client.List", err)
		return err
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			fmt.Printf("Normal Client 接收完成\n")
			break
		}
		if err != nil {
			fmt.Println("读取err: ", err)
			return err
		}
		fmt.Printf("Normal Client 从 Stream Server 接收到结果:name[%v],value[%v]\n", res.Pt.Name, res.Pt.Value)
	}
	return nil
}

//客户端发送流式RPC请求，服务器一次返回
func ClientSideStreamDemo(client pb.StreamServiceClient) error {
	stream, err := client.ClientSideStream(context.Background())
	if err != nil {
		return err
	}
	fmt.Printf("Stream Client 请求 Normal Server\n")
	for i := 1; i < 10; i++ {
		req := &pb.StreamReq{Pt: &pb.StreamPoint{Name: "客户端流式发送", Value: int32(i)}}
		fmt.Printf("Stream Client 发送请求: name[%v], value[%v], 发送完休息1s\n", req.Pt.Name, req.Pt.Value)
		err := stream.Send(req)
		if err != nil {
			return err
		}
		time.Sleep(time.Second)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}
	fmt.Printf("Stream Client 从 Normal Server接收到：name[%v], value[%v]\n", res.Pt.Name, res.Pt.Value)
	return nil
}

func DoubleSideStreamDemo(client pb.StreamServiceClient) error {
	stream, err := client.DoubleSideStream(context.Background())
	if err != nil {
		return err
	}

	for i := 0; i < 10; i++ {

		req := &pb.StreamReq{Pt: &pb.StreamPoint{Name: "stream client请求", Value: int32(i)}}
		fmt.Printf("client 发送：name:[%v], value[%v]\n", req.Pt.Name, req.Pt.Value)
		err = stream.Send(req)
		if err != nil {
			return nil
		}
		res, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		fmt.Printf("client 接收：name:[%v], value[%v]\n", res.Pt.Name, res.Pt.Value)

	}
	return nil
}
