package datarace

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestStringDataRace(t *testing.T)  {
	var str string
	go func(str1 string) {

		i := 0
		for {
			if i&1 == 0 {
				fmt.Println("奇数")
				str1 = "hello"
			}else {
				fmt.Println("偶数")
				str1 = "hello world"
			}
			i++
			time.Sleep(time.Second)
		}
	}(str)
	for  {
		fmt.Printf("%d, %s \n", len(str),str)
		time.Sleep(time.Second)
	}
}

var Counter int = 0
var Wait sync.WaitGroup

func TestIntDataRace(t *testing.T) {
	for routine := 1; routine <= 2; routine++ {
		Wait.Add(1)
		go Routine(routine)
	}
	Wait.Wait()
	fmt.Printf("Final Counter: %d\n", Counter)
}

func Routine(id int) {

	for count := 0; count < 2; count++ {
		value := Counter
		time.Sleep(1 * time.Second)
		value++
		Counter = value
	}
	Wait.Done()
}
