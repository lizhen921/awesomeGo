
## 安装delve

安装delve：

    go get -u github.com/go-delve/delve/cmd/dlv

或者（$GOPATH要提前设置好）

    $ git clone https://github.com/go-delve/delve.git $GOPATH/src/github.com/go-delve/delve
    $ cd $GOPATH/src/github.com/go-delve/delve
    $ make install


打包命令：

    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -gcflags "all=-N -l" -ldflags '-extldflags "-static"' -o main ./cmd/serve/main.go

参数解释：

    -a
    -gcflags
    -ldflags
    -o
    




