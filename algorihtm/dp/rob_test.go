package dp

import (
	"fmt"
	"testing"
)

func TestRob(t *testing.T) {
	nums := []int{4, 3}
	sum := rob(nums)
	fmt.Println(sum)
}
