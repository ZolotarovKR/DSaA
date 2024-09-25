package bin_coef

var memo [][]int = [][]int{{1}}

func BinCoef(r, c int) int {
	for j := len(memo); j <= r; j++ {
		memo = append(memo, make([]int, j+1))
		memo[j][0] = 1
		memo[j][j] = 1
		for i := 1; i < j; i++ {
			memo[j][i] = memo[j-1][i-1] + memo[j-1][i]
		}
	}
	return memo[r][c]
}
