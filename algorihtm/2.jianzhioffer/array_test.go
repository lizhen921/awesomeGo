package __jianzhioffer

import (
	"fmt"
	"testing"
)

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
