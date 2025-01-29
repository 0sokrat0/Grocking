package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	fmt.Println("Sum:", sum(arr, 0))                          // Сумма элементов
	fmt.Println("Max:", max(arr, 0, arr[0]))                  // Максимальный элемент
	fmt.Println("Binary Search index:", BinarySearch(arr, 6)) // Бинарный поиск
}

func sum(arr []int, index int) int {

	if index >= len(arr) {
		return 0
	}

	return arr[index] + sum(arr, index+1)

}

func max(arr []int, index int, currentMax int) int {
	if index >= len(arr) {
		return currentMax
	}
	if arr[index] > currentMax {
		currentMax = arr[index]
	}

	return max(arr, index+1, currentMax)

}

func BinarySearch(arr []int, target int) int {
	return binarySearchRecursive(arr, target, 0, len(arr)-1)
}

func binarySearchRecursive(arr []int, target int, low int, high int) int {
	if low > high {
		return -1
	}
	mid := low + (high-low)/2
	if arr[mid] == target {
		return mid
	}
	if arr[mid] > target {
		return binarySearchRecursive(arr, target, low, mid-1)
	}
	return binarySearchRecursive(arr, target, mid+1, high)
}

func search(nums []int, target int) int {

	if len(nums) == 0 {
		return -1
	}

	low, high := 0, len(nums)-1
	for low <= high {

		mid := low + (high-low)/2

		if nums[mid] == target {
			return mid
		}

		if nums[low] <= nums[mid] {
			if nums[low] <= target && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {

			if nums[mid] < target && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}
