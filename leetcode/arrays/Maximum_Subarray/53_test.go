package maximumsubarray

import (
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{-2, 1}, 1},
	}

	for _, test := range tests {
		actual := maxSubArray(test.nums)
		if actual != test.expected {
			t.Errorf("For nums %v, expected %d but got %d", test.nums, test.expected, actual)
		}

	}

}
