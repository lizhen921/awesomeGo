package __jianzhioffer

import (
	"sort"
	"testing"
)

//摆动序列
func wiggleMaxLength(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	count := 1
	pre, cur := 0, 0

	for i := 0; i < len(nums)-1; i++ {
		cur = nums[i+1] - nums[i]
		if (cur > 0 && pre <= 0) || (cur < 0 && pre >= 0) {
			count++
			pre = cur
		}

	}
	return count
}

func TestWiggleMaxLength(t *testing.T) {
	wiggleMaxLength([]int{3, 3, 3, 2, 5})

}

//s[j] (饼干)>= g[i]（学生）
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	count := 0
	j := len(s) - 1

	for i := len(g) - 1; i >= 0; i-- {
		if j >= 0 && s[j] >= g[i] {
			count++
			j--
		}
	}

	return count
}

//s[j] (饼干)>= g[i]（学生）
func findContentChildren2(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	count, j := 0, 0

	for i := 0; i < len(s); i++ {
		if j >= 0 && s[i] >= g[j] {
			count++
			j--
		}
	}

	return count
}