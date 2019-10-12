proto语法：
    定义的rpc接口必须有入参和出参

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


