package _defer

import (
	"fmt"
	"testing"
)

func TestDefer(t *testing.T) {
	fmt.Println(defer1())
	fmt.Println(defer2())
	fmt.Println(defer3())
}

func defer1() (r int) {
	defer func() {
		r++
	}()
	return r
}
func defer2() (r int) {
	t := 0
	defer func() {
		t++
	}()
	return t
}
func defer3() (r int) {
	defer func(r int) {
		r++
	}(r)
	return r
}