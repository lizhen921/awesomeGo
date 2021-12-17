package util

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	now := time.Now().UnixNano()
	fmt.Println(now)

	var t1 int64 = 1639554475000000000
	var t2 int64 = 1639727275000000000

	in := IsInNDay(2, t1, t2)

	fmt.Println(in)
}
