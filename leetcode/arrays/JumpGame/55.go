package main

func canJump(nums []int) bool {
	finalpos := len(nums) - 1
	for idx := len(nums) - 2; idx >= 0; idx-- {
		if idx+nums[idx] >= finalpos {
			finalpos = idx
		}

	}
	return finalpos == 0
}

// func jump(nums []int, leight int, start int) {
// 	for _, start := range nums {
// 		if start == 0 {
// 			return false
// 		}
// 	}
// }
