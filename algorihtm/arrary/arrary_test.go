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
