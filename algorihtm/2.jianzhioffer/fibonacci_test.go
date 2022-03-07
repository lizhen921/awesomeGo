package __jianzhioffer

func Fibonacci(n int) (res int) {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return Fibonacci(n) + Fibonacci(n-1)
}

func Fibonacci01(n int) (res int) {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	f1 := 0
	f2 := 1
	for i := 2; i <= n; i++ {
		res = f1 + f2
		f1 = f2
		f2 = res
	}

	return res
}
