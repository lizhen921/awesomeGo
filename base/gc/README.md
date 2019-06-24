# go垃圾回收

## 逃逸分析

在当前目录执行,`go build -gcflags=-m`，`-m`可以`print optimization decisions`，即打印编译过程中的优化内容，如下：

（使用`go tool compile -help`查看完整的参数及含义）

```
➜  gc git:(master) ✗ go build -gcflags=-m
# awesomeGo/base/gc
./escape.go:27:6: can inline Center                             //27行，方法Center可以作为内联函数
./escape.go:34:8: inlining call to Center                       //34行，将Center方法转为内联
./escape.go:10:17: Sum make([]int, 100) does not escape         //10行，Sum方法中make([]int, 100)，没有逃逸
./escape.go:27:13: Center c does not escape                     //27行，Center中的c对象没有逃逸
./escape.go:35:15: c.X escapes to heap                          //c.X逃逸至heap中
./escape.go:35:20: c.Y escapes to heap                          //c.Y逃逸至heap中
./escape.go:33:10: CenterCoursor new(Coursor) does not escape   //new(Coursor)没有逃逸
./escape.go:35:13: CenterCoursor ... argument does not escape   //35行参数没有逃逸
```

go引入了逃逸分析，即如果一个对象没有在当前方法之外的地方被引用，则可以称该方法没有逃逸，即该对象会被分配到stack中；否则的话，会分配到堆中。