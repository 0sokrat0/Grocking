package kthlargestelementinanarray

func findKthLargest(nums []int, k int) int {
	left, right := 0, len(nums)-1
	targetIndex := len(nums) - k

	for left <= right {
		low, high := partition(nums, left, right)

		if targetIndex >= low && targetIndex <= high {
			return nums[targetIndex]
		} else if targetIndex < low {
			right = low - 1
		} else {
			left = high + 1
		}
	}
	return -1
}

func partition(nums []int, left, right int) (int, int) {
	/*************  ✨ Codeium Command ⭐  *************/
	// partition rearranges the slice nums such that all elements less than the pivot are on the left
	// of the pivot, all elements greater than the pivot are on the right, and all elements equal to the
	// pivot are in the middle. It returns the start and end indices of the equal elements.
	/******  3d574cc3-b2e3-424e-a8ab-7f5b804ce62c  *******/
	pivot := nums[right]
	i, j, k := left, left, right

	for j <= k {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
		} else if nums[j] > pivot {
			nums[j], nums[k] = nums[k], nums[j]
			k--
		} else {
			j++
		}
	}

	return i, k
}
