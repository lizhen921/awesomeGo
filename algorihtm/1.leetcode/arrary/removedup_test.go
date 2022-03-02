package arrary

import (
	"fmt"
	"testing"
)

func TestRemoveDup(t *testing.T) {
	nums := []int{1, 2, 3, 3, 3, 4, 5, 5}
	len := removeDuplicates(nums)
	fmt.Println(nums)
	fmt.Println(len)
	for i := 0; i < len; i++ {
		print(nums[i])
	}
}
