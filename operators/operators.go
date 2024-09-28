package operators

func maxValue(ns []int) int {
	d := make([][]int, len(ns))
	for i := range d {
		d[i] = make([]int, i+1)
		d[i][i] = ns[i]
	}
	for j := 1; j < len(d); j++ {
		for i := j - 1; i >= 0; i-- {
			for s := i; s < j; s++ {
				t := d[j][s+1] + d[s][i]
				if t > d[j][i] {
					d[j][i] = t
				}
				t = d[j][s+1] * d[s][i]
				if t > d[j][i] {
					d[j][i] = t
				}
			}
		}
	}
	return d[len(d)-1][0]
}
