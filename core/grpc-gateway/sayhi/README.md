生成grpc服务文件：
`protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. ./helloworld.proto `
生成grpc-gateway文件：
`protoc -I/usr/local/include -I. -I$GOPATH/src  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. ./helloworld.proto`

先启动grpc服务(grpc-server/main.go)，再启动grpc-gateway服务(grpc-gw-server/main.go)，然后测试
`curl -X POST -k http://localhost:8081/v1/example/hello -d '{"name": " world"}'`
返回：
`{"message":"Hello  world"}`



protoc -I. -I$GOPATH/src --go_out=plugins=grpc:. ./com/autohome/summary/common/*

found several packages [com_autohome_summary_common_proto, com_autohome_summary_common] in 
/Users/lizhen/go/src/git.corpautohome.com/autorcm/rcm/vendor/git.corpautohome.com/autorcm/PJ_RECOMMEND_SUMMARY/power-summary-common/src/main/proto/com/autohome/summary/common;
/Users/lizhen/go/src/git.corpautohome.com/autorcm/rcm/vendor/git.corpautohome.com/autorcm/PJ_RECOMMEND_SUMMARY/power-summary-common/src/main/proto/com/autohome/summary/common' less... (⌘F1) 
Inspection info: Reports invalid imports.
