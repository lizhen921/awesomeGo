package goroutine

import (
	"fmt"
	"testing"
)

func TestGor(t *testing.T) {
	m := make(map[string]int)
	s := make([]string, 0, 1)
	a := s[1]
	fmt.Println(a)
	for i := 0; i < 200; i++ {
		go func(n int) {
			m[fmt.Sprintf("%d", n)] = n
		}(i)
	}
}
