package main

import "fmt"

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))

}

func maxArea(height []int) int {
	max := 0
	right := len(height) - 1

	for left := 0; left < right; {
		width := right - left
		h := min(height[right], height[left])
		if height[left] < height[right] {
			left++
		} else if height[left] >= height[right] {
			right--
		}
		if width*h > max {
			max = width * h
		}
	}
	return max

}

// [1,8,6,2,5,4,8,3,7]

// func mini(x int, y int) int {
// 	if x > y {
// 		return y
// 	} else {
// 		return x
// 	}
// }
