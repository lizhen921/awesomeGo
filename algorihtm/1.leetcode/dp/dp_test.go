package dp

import (
	"fmt"
	"testing"
)

func TestFib(t *testing.T) {
	fmt.Println(fib(5))
}

func TestClimbStairs01(t *testing.T) {
	fmt.Println(climbStairs01(5))
}

func TestClimbStairs02(t *testing.T) {
	fmt.Println(climbStairs02([]int{10, 15, 20}))
}
