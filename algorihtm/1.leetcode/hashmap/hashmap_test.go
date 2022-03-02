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

/*
给你四个整数数组 nums1、nums2、nums3 和 nums4 ，数组长度都是 n ，请你计算有多少个元组 (i, j, k, l) 能满足：

- 0 <= i, j, k, l < n
- nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0

示例 1：

输入：nums1 = [1,2], nums2 = [-2,-1], nums3 = [-1,2], nums4 = [0,2]
输出：2
解释：
两个元组如下：
1. (0, 0, 0, 1) -> nums1[0] + nums2[0] + nums3[0] + nums4[1] = 1 + (-2) + (-1) + 2 = 0
2. (1, 1, 0, 0) -> nums1[1] + nums2[1] + nums3[0] + nums4[0] = 2 + (-1) + (-1) + 0 = 0

示例 2：

输入：nums1 = [0], nums2 = [0], nums3 = [0], nums4 = [0]
输出：1

*/
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	sumMap := make(map[int]int)
	for _, i := range nums1 {
		for _, j := range nums2 {
			sumMap[i+j]++
		}
	}
	count := 0
	for _, k := range nums3 {
		for _, l := range nums4 {
			if _, ok := sumMap[-(k + l)]; ok {
				count += sumMap[-(k + l)]
			}
		}
	}
	return count
}
func TestFourSumCount(t *testing.T) {
	num1 := []int{-1, -1}
	num2 := []int{-1, 1}
	num3 := []int{-1, 1}
	num4 := []int{1, -1}
	fmt.Println(fourSumCount(num1, num2, num3, num4))
}

/*
给定一个整数数组nums和一个整数目标值target，请你在该数组中找出 和为目标值target的那两个整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案。


示例 1：

输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。


示例 2：

输入：nums = [3,2,4], target = 6
输出：[1,2]


示例 3：

输入：nums = [3,3], target = 6
输出：[0,1]

*/
func twoSum(nums []int, target int) []int {

	numMap := make(map[int]int)
	for i, v := range nums {
		if p, ok := numMap[target-v]; ok {
			return []int{p, i}
		}
		numMap[v] = i
	}

	return nil
}

func TestTwoSum(t *testing.T) {

	n := twoSum([]int{3, 3, 3}, 6)
	fmt.Println(n)
}

/*
为了不在赎金信中暴露字迹，从杂志上搜索各个需要的字母，组成单词来表达意思。

给你一个赎金信 (ransomNote) 字符串和一个杂志(magazine)字符串，判断 ransomNote 能不能由 magazines 里面的字符构成。

如果可以构成，返回 true ；否则返回 false 。

magazine 中的每个字符只能在 ransomNote 中使用一次。

示例 1：

输入：ransomNote = "a", magazine = "b"
输出：false

*/
func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	magMap := make(map[rune]int)
	for _, v := range magazine {
		magMap[v]++
	}

	for _, v := range ransomNote {
		if m, ok := magMap[v]; ok && m > 0 {
			magMap[v]--
		} else {
			return false
		}
	}
	return true
}

/*
给你一个包含 n 个整数的数组nums，判断nums中是否存在三个元素 a，b，c ，使得a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。

示例 1：

输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
示例 2：

输入：nums = []
输出：[]
示例 3：

输入：nums = [0]
输出：[]

*/

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0, 0)
	for i := 0; i < len(nums)-2; i++ {
		n1 := nums[i]
		if n1 > 0 {
			break
		}
		if i > 0 && n1 == nums[i-1] {
			continue
		}
		spoint := i + 1
		epoint := len(nums) - 1
		for spoint < epoint {
			n2 := nums[spoint]
			n3 := nums[epoint]
			if n1+n2+n3 == 0 {
				res = append(res, []int{n1, n2, n3})
				for spoint < epoint && nums[spoint] == n2 {
					spoint++
				}
				for spoint < epoint && nums[epoint] == n3 {
					epoint--
				}
			} else if n1+n2+n3 > 0 {
				epoint--
			} else {
				spoint++
			}
		}
	}
	return res
}

/*
给你一个由 n 个整数组成的数组nums ，和一个目标值 target 。
请你找出并返回满足下述全部条件且不重复的四元组[nums[a], nums[b], nums[c], nums[d]]（若两个四元组元素一一对应，则认为两个四元组重复）：

0 <= a, b, c, d< n
a、b、c 和 d 互不相同
nums[a] + nums[b] + nums[c] + nums[d] == target
你可以按 任意顺序 返回答案 。

示例 1：

输入：nums = [1,0,-1,0,-2,2], target = 0
输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
示例 2：

输入：nums = [2,2,2,2,2], target = 8
输出：[[2,2,2,2]]

*/

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0, 0)
	for i := 0; i < len(nums)-3; i++ {
		n1 := nums[i]

		if i > 0 && n1 == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			n2 := nums[j]

			if j > i+1 && n2 == nums[j-1] {
				continue
			}
			start := j + 1
			end := len(nums) - 1
			for start < end {
				n3 := nums[start]
				n4 := nums[end]
				sum := n1 + n2 + n3 + n4
				if sum == target {
					res = append(res, []int{n1, n2, n3, n4})
					for start < end && nums[start] == n3 {
						start++
					}
					for start < end && nums[end] == n4 {
						end--
					}
				} else if sum < target {
					start++
				} else {
					end--
				}
			}

		}
	}
	return res
}
