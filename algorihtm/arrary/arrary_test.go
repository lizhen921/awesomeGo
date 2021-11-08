package arrary

import (
	"fmt"
	"testing"
)

//二分查找
/*
给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
比如：
输入: nums = [-1,0,3,5,9,12], target = 9
输出: 4
解释: 9 出现在 nums 中并且下标为 4


输入: nums = [-1,0,3,5,9,12], target = 2
输出: -1
解释: 2 不存在 nums 中因此返回 -1

https://leetcode-cn.com/problems/binary-search/
*/
func BinarySearch(slice []int, target int) int {
	left := 0
	right := len(slice) - 1

	for left <= right {
		mid := left + (right-left)/2
		if target < slice[mid] {
			right = mid - 1
		} else if target > slice[mid] {
			left = mid + 1
		} else {
			return mid
		}
	}

	return -1
}

func TestBinarySearch(t *testing.T) {
	sortedSlice := []int{-1, 0, 3, 5, 9, 12}
	index := BinarySearch(sortedSlice, 9)
	fmt.Println(index)
}

/*
题目：https://leetcode-cn.com/problems/remove-element/

给你一个数组 nums和一个值val，你需要原地移除所有数值等于val的元素，并返回移除后数组的新长度。

不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。

元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

比如：

输入：nums = [3,2,2,3], val = 3

输出：2, nums = [2,2]

解释：函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。你不需要考虑数组中超出新长度后面的元素。例如，函数返回的新长度为 2 ，而 nums = [2,2,3,3] 或 nums = [2,2,0,0]，也会被视作正确答案。

输入：nums = [0,1,2,2,3,0,4,2], val = 2

输出：5, nums = [0,1,4,0,3]

解释：函数应该返回新的长度 5, 并且 nums 中的前五个元素为 0, 1, 3, 0, 4。注意这五个元素可为任意顺序。你不需要考虑数组中超出新长度后面的元素。

*/

func removeElement(nums []int, val int) int {
	start := 0

	for i, v := range nums {
		if v != val {
			nums[start] = nums[i]
			start++
		}
	}
	return start
}

func TestRemoveElement(t *testing.T) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	numsLen := removeElement(nums, 2)
	fmt.Println(numsLen, nums)
}

/*
给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。

示例 1： 输入：nums = [-4,-1,0,3,10] 输出：[0,1,9,16,100] 解释：平方后，数组变为 [16,1,0,9,100]，排序后，数组变为 [0,1,9,16,100]

示例 2： 输入：nums = [-7,-3,2,3,11] 输出：[4,9,9,49,121]

*/

func sortedSquares(nums []int) []int {
	numsLen := len(nums)
	ans := make([]int, numsLen, numsLen)

	a, b := 0, numsLen-1
	for i := range ans {
		aa, bb := nums[a]*nums[a], nums[b]*nums[b]

		if aa >= bb {
			ans[numsLen-1-i] = aa
			a++
		} else {
			ans[numsLen-1-i] = bb
			b--
		}
	}
	return ans
}

func TestSortedSquares(t *testing.T) {
	nums := []int{-10, -3, 0, 3, 5, 12, 15}
	fmt.Println(sortedSquares(nums))
}

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
