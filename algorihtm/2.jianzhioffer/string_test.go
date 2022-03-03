package __jianzhioffer

import (
	"fmt"
	"testing"
)

func RepalceStr(s string) string {
	strB := make([]byte, 0, len(s))
	for i := range s {
		if s[i] == ' ' {
			strB = append(strB, []byte("%20")...)
		} else {
			strB = append(strB, s[i])
		}
	}
	return string(strB)
}

func TestRepalceStr(t *testing.T) {
	fmt.Println(RepalceStr("we are happy"))
}
