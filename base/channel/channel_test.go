package channel

import (
	"testing"
	"time"
)

var c = make(chan int, 10)
var a string

func f() {
	c <- 0
	time.Sleep(time.Second * 3)
	a = "hello, world"

}

func TestChanne(t *testing.T) {
	go f()
	<-c
	print(a)
}
