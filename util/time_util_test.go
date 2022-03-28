package util

import (
	"fmt"
	"testing"
)

func TestTime(t *testing.T) {
	//now := time.Now().UnixNano()
	fmt.Println(TimestampToTimeStr(1647878400))
}
