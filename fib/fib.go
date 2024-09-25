package fib

var memo []int = []int{0, 1}

func Fib(n int) int {
	for i := len(memo); i <= n; i++ {
		memo = append(memo, memo[i-2]+memo[i-1])
	}
	return memo[n]
}
