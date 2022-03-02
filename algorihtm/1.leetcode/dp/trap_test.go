package dp

import (
	"fmt"
	"testing"
)

func TestTrap(t *testing.T) {
	height := []int{0, 1, 10, 0, 2, 1, 0, 1, 3, 2, 8, 1, 2, 1}
	sum := trap(height)
	fmt.Println(sum)
	sum2 := trap2(height)
	fmt.Println(sum2)
}
