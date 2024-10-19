package backpack

import "testing"

func TestBackpack(t *testing.T) {
	testCases := []struct {
		items              []struct{ weight, price int }
		capacity, expected int
	}{
		{
			[]struct{ weight, price int }{
				{3, 1},
				{4, 6},
				{5, 4},
				{8, 7},
				{9, 6},
			},
			13, 13,
		},
	}

	for _, tc := range testCases {
		actual := Backpack(tc.items, tc.capacity)
		if actual != tc.expected {
			t.Logf("items: %v", tc.items)
			t.Logf("capacity: %v", tc.capacity)
			t.Fatalf(
				"Expected total price is %v, but actual %v",
				tc.expected,
				actual,
			)
		}
	}
}
