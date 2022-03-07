package __jianzhioffer

import (
	"fmt"
	"testing"
)

func search(arr []int, val int) (found bool, index int) {
	i := 0
	j := len(arr) - 1
	for i <= j {
		mid := i + (j-i)/2
		midV := arr[mid]
		if midV > val {
			j = mid - 1
		} else if midV < val {
			i = mid + 1
		} else {
			return true, i + (j-i)/2
		}
	}
	return false, 0
}

// [3,4,5,1,2] 为 [1,2,3,4,5]

//[0,1,4,4,5,6,7]
//[4,5,6,7,0,1,4]
//[0,1,4,4,5,6,7]

func minArray(numbers []int) int {
	i := 0
	j := len(numbers) - 1

	for i < j {
		mid := i + (j-i)/2
		if numbers[mid] < numbers[j] {
			//在前
			j = mid
		} else if numbers[mid] > numbers[j] {
			//在后
			i = mid + 1
		} else {
			//相等，迁移一位，再判断
			j--
		}
	}
	return numbers[i]
}

func minArray01(number []int) int {
	i := 0
	j := len(number) - 1

	if number[i] < number[j] {
		return number[i]
	}

	for i < j {
		if number[i] < number[j] {
			return number[i]
		}
		mid := i + (j-i)/2
		if number[i] < number[mid] {
			i = mid
		} else if number[i] > number[mid] {
			j = mid
		} else {
			i++
		}
	}
	return number[i]
}

func TestMinArray(t *testing.T) {

}

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	i := 0
	j := len(matrix[0]) - 1

	for i <= len(matrix)-1 && j >= 0 {
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] > target {
			j--
			continue
		}
		if matrix[i][j] < target {
			i++
			continue
		}
	}
	return false
}

//判断某个值，是否在有序的二维数组中
func TestIsInArrary(t *testing.T) {
	a1 := [][]int{
		[]int{1, 2, 8, 9},
		[]int{2, 5, 9, 12},
		[]int{4, 7, 10, 13},
		[]int{6, 8, 11, 15},
	}
	fmt.Println(findNumberIn2DArray(a1, 10))
}
