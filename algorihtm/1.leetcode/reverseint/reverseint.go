package reverseint

/*
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

注意：

假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-integer
*/
func reverse(x int32) int32 {
	var res int32
	for x != 0 {
		restemp := res*10 + x%10
		if (int32(restemp)-x%10)/10 != res { //判断越界
			return 0
		}
		res = restemp
		x = x / 10
	}

	return res
}
