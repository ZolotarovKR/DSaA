package rod_cutting

func CutRod(length int, prices []int) int {
	ps := make([]int, length+1)
	for i := 1; i <= length; i++ {
		for j := 1; j <= i; j++ {
			p := prices[j] + ps[i-j]
			if p > ps[i] {
				ps[i] = p
			}
		}
	}
	return ps[length]
}
