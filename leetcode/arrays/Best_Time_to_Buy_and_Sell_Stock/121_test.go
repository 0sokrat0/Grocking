package besttimetobuyandsellstock

import "testing"

func maxProfitTest(t *testing.T) {
	tests := []struct {
		prices   []int
		expected int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{1, 2, 3, 4, 5}, 4},
		{[]int{3, 2, 6, 1, 3}, 4},
	}
	for _, test := range tests {
		actual := maxProfit(test.prices)
		if actual != test.expected {
			t.Errorf("For prices %v, expected %d but got %d", test.prices, test.expected, actual)
		}

	}

}
