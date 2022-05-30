package util

import (
	"fmt"
	"testing"
)

func TestTime(t *testing.T) {

	s := fmt.Sprintf("%d%s的奥迪A6L用户", 1, "%")
	fmt.Println(s)
	//now := time.Now().UnixNano()
	fmt.Println(TimestampToTimeStr(1647878400))
}
