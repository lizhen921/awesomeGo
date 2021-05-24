package arrary

import (
	"fmt"
	"testing"
)

func TestArrary(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("%v\n", &arr)
	for i, _ := range arr {
		fmt.Printf("%v\n", &arr[i])
	}

	var m map[string]string
	v, ok := m["1"]
	fmt.Printf("%v, %v", v, ok)
	m["1"] = "1"
}
