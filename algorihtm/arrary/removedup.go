package arrary

/*
给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。

不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array
*/
func removeDuplicates(nums []int) int {

	if len(nums) <= 1 {
		return 1
	}

	index := 1
	for i, _ := range nums {
		if i > 0 && nums[i] != nums[i-1] {
			nums[index] = nums[i]
			index++
		}
	}

	nums = nums[:index]

	return index
}
