gRPC 是一个高性能、开源和通用的 RPC 框架，面向服务端和移动端，基于 HTTP/2 设计。

proto语法：
    定义的rpc接口必须有入参和出参

## GRPC入门Demo

1. 创建hello.proto文件，定义HelloService接口：
```$xslt
syntax = "proto3";

package helloworld;

message Req{
    string value = 1;
}
message Res{
    string value = 1;
}

service HelloService{
    rpc Hello (Req) returns (Res);
}

```

2. 使用protoc-gen-go内置的gRPC插件生成gRPC代码：

```$xslt
$ protoc --go_out=plugins=grpc:. hello.proto
```

3. 查看`hello.pb.go`，gRPC插件会为服务端和客户端生成不同的接口：

```$xslt
// HelloServiceServer is the server API for HelloService service.
type HelloServiceServer interface {
	Hello(context.Context, *Req) (*Res, error)
}
```
```$xslt
// HelloServiceClient is the client API for HelloService service.
type HelloServiceClient interface {
	Hello(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Res, error)
}

```
gRPC通过context.Context参数，为每个方法调用提供了上下文支持。客户端在调用方法的时候，可以通过可选的grpc.CallOption类型的参数提供额外的上下文信息。

4. 编写server目录`hello.go`，实现grpc服务端
```$xslt
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

```

5. 编写client目录下`hello.go`，实现客户端调用
```$xslt
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

	client := helloworld.NewHelloServiceClient(conn)  //grpc会为每个service，生成一个NewXXXClient方法
	res, err := client.Hello(context.Background(), &helloworld.Req{Value: "client: helloworld"})
	if err != nil {
		fmt.Printf("请求err： %v", err)
	}
	fmt.Println(res)
}

```

## 普通RPC 
`参看Demo: search.proto及相关方法`
proto中声明到RPC接口，生成pb.go之后，会生成相应到client和server方法。

server端：需要实现pb.go中的XXXServer接口，然后将自己实现的对象调用RegisterXXXServer方法，注册rpc服务；

client端：需要调用NewXXXClient来获取rpc客户端，然后直接调用pb.go中的XXXClient接口的方法即可，gprc会直接寻找对应地址注册的server方法。

## 流式RPC
`参看Demo: stream.proto及相关方法`
除了普通rpc接口，还支持流式rpc接口。

流式gRPC，分为三种类型：

    Server-side streaming RPC：服务器端流式 RPC
    Client-side streaming RPC：客户端流式 RPC
    Bidirectional streaming RPC：双向流式 RPC

声明方式见demo，希望哪一方（client/server）使用流式发送数据，就在哪边参数(res/req)前增加`stream`关键字

server端，同样要实现接口来提供自己的RPC服务，其中一个的地方就是将普通req/res类型转变为stream类型。

每个stream对象都有send和recv方法，用来发送或者接收参数。

流式RPC与普通RPC主要区别在与req或者res，是使用普通对象传递还是通过stream方式。普通的RPC直接使用参数或者返回值就可以了，而流式RPC则需要send/recv方法才可以。


