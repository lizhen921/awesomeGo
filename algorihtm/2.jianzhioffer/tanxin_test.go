package __jianzhioffer

import (
	"sort"
)

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
