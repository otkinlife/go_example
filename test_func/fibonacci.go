package test_func

// Fibonacci 斐波那契数列
func Fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = Fibonacci(n-1) + Fibonacci(n-2)
	}
	return res
}
