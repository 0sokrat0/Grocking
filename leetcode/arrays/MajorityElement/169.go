package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 2, 3}
	fmt.Print(majorityElement(nums))

}

// func majorityElement(nums []int) int {
// 	m := make(map[int]int)
// 	MaxKey := 0
// 	MaxValue := 0

// 	for _, v := range nums {
// 		m[v]++
// 		fmt.Println(m)
// 		if m[v] > MaxValue {
// 			MaxValue = m[v]
// 			MaxKey = v
// 		}
// 	}
// 	return MaxKey
// }

func majorityElement(nums []int) int {
	count := 0
	candidate := 0

	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}
