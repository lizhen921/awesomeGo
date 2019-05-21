生成grpc服务文件：
`protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. ./helloworld.proto `
生成grpc-gateway文件：
`protoc -I/usr/local/include -I. -I$GOPATH/src  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. ./helloworld.proto`

先启动grpc服务，再启动grpc-gateway服务，然后测试
`curl -X POST -k http://localhost:8081/v1/example/hello -d '{"name": " world"}'`
返回：
`{"message":"Hello  world"}`