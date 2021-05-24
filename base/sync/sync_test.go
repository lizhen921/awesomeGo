package sync

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var a string
var once sync.Once

func setup() {
	println("123")
	a = "hello, world"
}

func doprint() {
	once.Do(setup)
	fmt.Println(a)
}

func twoprint() {
	go doprint()
	go doprint()
	time.Sleep(time.Second * 5)
}

func TestOnce(t *testing.T) {
	twoprint()
}
