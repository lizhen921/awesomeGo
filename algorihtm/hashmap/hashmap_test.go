package hashmap

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

/*
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。

输入: s = "anagram", t = "nagaram"
输出: true

输入: s = "rat", t = "car"
输出: false
*/

func isAnagram(s string, t string) bool {
	sliceS := strings.Split(s, "")
	sliceT := strings.Split(t, "")
	sort.Strings(sliceS)
	sort.Strings(sliceT)
	s = strings.Join(sliceS, "")
	t = strings.Join(sliceT, "")
	if s == t {
		return true
	}
	return false
}

func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	strMap := make(map[rune]int)
	for _, v := range s {
		strMap[v]++
	}

	for _, v := range t {
		strMap[v]--
		if strMap[v] < 0 {
			return false
		}
	}
	return true
}

/*
给定两个数组，编写一个函数来计算它们的交集。

示例 1：

输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2]

示例 2：

输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[9,4]
*/

func intersection(nums1 []int, nums2 []int) []int {
	numMap := make(map[int]int)
	num := make([]int, 0, 0)
	for _, v := range nums1 {
		numMap[v] = 1
	}

	for _, v := range nums2 {
		if n, ok := numMap[v]; ok && n == 1 {
			numMap[v]--
			num = append(num, v)

		}
	}
	return num
}

func intersection2(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	num := make([]int, 0, 0)

	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] == nums2[j] {
			if len(num) == 0 || num[len(num)-1] != nums1[i] {
				num = append(num, nums1[i])
			}
			i++
			j++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			i++
		}

	}
	return num
}

/*
编写一个算法来判断一个数 n 是不是快乐数。

「快乐数」定义为：

对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
如果 可以变为  1，那么这个数就是快乐数。
如果 n 是快乐数就返回 true ；不是，则返回 false 。

示例 1：

输入：n = 19
输出：true
解释：
12 + 92 = 82
82 + 22 = 68
62 + 82 = 100
12 + 02 + 02 = 1
示例 2：

输入：n = 2
输出：false

*/
func isHappy(n int) bool {
	nMap := make(map[int]bool)
	for n != 1 {
		n = getNext(n)
		if nMap[n] {
			return false
		}
		nMap[n] = true
	}
	return true
}

func getNext(n int) int {
	var res int
	for n > 0 {
		res += (n % 10) * (n % 10)
		n = n / 10
	}
	return res
}

func TestIsHappy(t *testing.T) {
	fmt.Println(isHappy(19))
}
