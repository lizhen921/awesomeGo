package main

import (
	pb "awesomeGo/core/grpc/go-grpc-example/proto"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"
)

type StreamService struct {
}

func (s *StreamService) ServerSideStream(req *pb.StreamReq, stream pb.StreamService_ServerSideStreamServer) error {
	fmt.Printf("服务器端流式 RPC接口 接收到请求，name[%v], value[%v]\n", req.Pt.Name, req.Pt.Value)
	for i := 1; i < 10; i++ {
		res := &pb.StreamRes{
			Pt: &pb.StreamPoint{
				Name:  "服务器端流式RPC结果",
				Value: int32(i),
			},
		}
		fmt.Printf("stream 服务端返回 name[%v], value[%v], 发送完休息1s\n", res.Pt.Name, res.Pt.Value)
		err := stream.Send(res)
		if err != nil {
			fmt.Println("服务器端流式 RPC err:", err)
			return err
		}
		time.Sleep(time.Second)
	}
	return nil
}

func (s *StreamService) ClientSideStream(stream pb.StreamService_ClientSideStreamServer) error {
	fmt.Println("Normal Client 接收到 Stream Client 请求")
	resName := ""
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			res := &pb.StreamRes{Pt: &pb.StreamPoint{Name: "全部结果" + resName, Value: 99}}
			fmt.Printf("全部接收完成，返回结果[%v][%v]\n", res.Pt.Name, res.Pt.Value)
			return stream.SendAndClose(res)
		}
		if err != nil {
			return err
		}
		resName = resName + "-" + strconv.FormatInt(int64(req.Pt.Value), 10)
		fmt.Printf("接收到请求: name[%v], value[%v]\n", req.Pt.Name, req.Pt.Value)

	}
	return nil
}

func (s *StreamService) DoubleSideStream(stream pb.StreamService_DoubleSideStreamServer) error {
	fmt.Println("双向流式 RPC 接收到请求")
	for {
		v := int32(rand.Int())
		res := &pb.StreamRes{Pt: &pb.StreamPoint{Name: "", Value: v}}
		err := stream.Send(res)
		if err != nil {
			return err
		}
		fmt.Printf("double stream server返回 name:[%v], value:[%v]\n", res.Pt.Name, res.Pt.Value)

		req, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("server端接收完成")
			return nil
		}
		if err != nil {
			return err
		}
		fmt.Printf("double stream server接收到 name:[%v], value:[%v]\n", req.Pt.Name, req.Pt.Value)

	}
	return nil
}

func main() {
	server := grpc.NewServer()
	pb.RegisterStreamServiceServer(server, &StreamService{})

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("启动server失败: %v", err)
	}
	server.Serve(lis)
}
