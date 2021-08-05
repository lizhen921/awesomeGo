package test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestCompare(t *testing.T) {
	s := VersionCompare("11.11.5", "11.1.0") > 0
	fmt.Println(s)
}

func VersionCompare(v1, v2 string) int {
	segs1 := strings.Split(v1, ".")
	segs2 := strings.Split(v2, ".")

	for i := 0; true; i++ {
		// 相等
		if len(segs1) <= i && len(segs2) <= i {
			return 0
		}

		// v1为空 v2不为空 小于
		if len(segs1) <= i {
			return -1
		}
		// v2为空 v1不为空 大于
		if len(segs2) <= i {
			return 1
		}

		// 不合法 -2
		n1, err1 := strconv.Atoi(segs1[i])
		if err1 != nil {
			return -2
		}
		n2, err2 := strconv.Atoi(segs2[i])
		if err2 != nil {
			return -2
		}

		if n1 > n2 {
			return 1
		} else if n1 < n2 {
			return -1
		}
	}
	return 0
}
