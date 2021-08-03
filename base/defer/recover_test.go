package _defer

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	defer catch("f")

	g()
}

func g() {
	defer m()

	panic("g panic")
}

func m() {
	panic("m panic")
}

func catch(funcname string) {
	if r := recover(); r != nil {
		fmt.Println(funcname, "recover:", r)
	}
}


