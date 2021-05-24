package string

import (
	"fmt"
	"testing"
)

func TestStringChar(t *testing.T) {
	testString := "中"

	//fmt.Println(testString[4])
	fmt.Println("---")
	for i := 0; i < len(testString); i++ { //只要使用下标索引，就是按byte取，go中用int8表示
		fmt.Println(testString[i])
	}

	fmt.Println("+++")
	for i, v := range testString { //只有for-range时，遍历得到的v，才是char。即go处理成的rune（int32）表示
		fmt.Println(testString[i])
		fmt.Println(testString[i+1])
		fmt.Println("---")
		fmt.Println(v)
	}

}
