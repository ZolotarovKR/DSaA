package backpack

func Backpack(items []struct{ weight, price int }, capacity int) int {
	m := make([][]int, len(items)+1)
	for j := range m {
		m[j] = make([]int, capacity+1)
	}
	for j := 1; j < len(m); j++ {
		for i := 1; i < len(m[j]); i++ {
			m[j][i] = m[j-1][i]
			if items[j-1].weight > i {
				continue
			}
			t := i - items[j-1].weight
			if m[j-1][t]+items[j-1].price > m[j-1][i] {
				m[j][i] = m[j-1][t] + items[j-1].price
			}
		}
	}
	return m[len(items)][capacity]
}
