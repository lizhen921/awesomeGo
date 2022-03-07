package __jianzhioffer

import (
	"fmt"
	"testing"
)

// 如果n为负数，则可能出现死循环，负数最高位为1，右移之后，高位仍然会补1，最终变成0xFFFFFFFF而陷入死循环
func NumberOf1(n int) (count int) {
	for n != 0 {
		if n&1 == 1 {
			count++
		}
		n = n >> 1
	}
	return n

}

//
func hammingWeight(num uint32) int {
	var i uint32 = 1
	count := 0
	for i != 0 {
		if i&num != 0 {
			count++
		}
		i = i << 1
	}
	return count
}

func hammingWeight02(num uint32) int {
	count := 0

	for num != 0 {
		num = num & (num - 1)
		count++
	}

	return count
}

func TestNumber02(t *testing.T) {
	fmt.Println(hammingWeight(00000000000000000000011100001011))
}
