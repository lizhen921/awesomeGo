package reverseint

import "fmt"

/*
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

你能不将整数转为字符串来解决这个问题吗？
*/
func isPalindrome(x int) bool {
	if x < 0 || (x > 0 && x%10 == 0) { //负数和个位为0的，不可能是回文数
		return false
	}
	temp := 0
	for x > temp {
		temp = temp*10 + x%10

		x = x / 10
	}
	fmt.Println("x: ", x)
	fmt.Println("temp: ", temp)
	return x == temp || x == temp/10
}
