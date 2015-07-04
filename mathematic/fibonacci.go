package mathematic

func Fibonacci(n int) int {
	if n == 1 || n == 0 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func FibonacciDyamic(n int) int {
	arr := make([]int, n+1, n+1)
	arr[0] = 0
	arr[1] = 1
	for x := 2; x <= n; x++ {
		arr[x] = -1
	}
	return fibonacciDyamic(n, arr)
}

func fibonacciDyamic(n int, arr []int) int {
	if val := arr[n]; val != -1 {
		return val
	}
	arr[n-1] = fibonacciDyamic(n-1, arr)
	arr[n-2] = fibonacciDyamic(n-2, arr)
	return arr[n-1] + arr[n-2]
}
