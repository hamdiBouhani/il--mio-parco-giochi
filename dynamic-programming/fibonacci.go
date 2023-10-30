package dynamicprogramming

type GenericMap[K comparable, V any] map[K]V


var memo = make(GenericMap[int, int])

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	if _, ok := memo[n]; !ok {
		memo[n] = Fibonacci(n-1) + Fibonacci(n-2)
	}

	return memo[n]
}