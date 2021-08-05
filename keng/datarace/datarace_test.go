package datarace

import (
	"fmt"
	"testing"
	"time"
)

func TestDataRace(t *testing.T)  {
	var str string
	go func(str1 string) {

		i := 0
		for {
			if i&1 == 0 {
				//fmt.Println("奇数")
				str1 = "hello"
			}else {
				//fmt.Println("偶数")
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