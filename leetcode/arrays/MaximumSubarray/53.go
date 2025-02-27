package maximumsubarray

func maxSubArray(nums []int) int {
	subarray := nums[0]
	sum := nums[0]

	for i := 1; i < len(nums); i++ {
		if sum < 0 {
			sum = nums[i]
		} else {
			sum += nums[i]
		}

		if sum > subarray {
			subarray = sum
		}
	}
	return subarray
}
