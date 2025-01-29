package main

import "fmt"

func main() {
	numbers := []int{2, 7, 11, 15}
	fmt.Println(twoSum(numbers, 9))

}

func twoSum(numbers []int, target int) []int {
	right := len(numbers) - 1
	res := []int{}

	for left := 0; left < right; {
		if numbers[right]+numbers[left] > target {
			right--
		} else if numbers[right]+numbers[left] < target {
			left++
		} else if numbers[right]+numbers[left] == target {
			res = append(res, left+1, right+1)
			return res
		}
	}
	return res
}
