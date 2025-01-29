package quicksort

func QuickSort(arr []int) []int {
	//Базовый случай
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)/2]
	left := []int{}
	right := []int{}
	equal := []int{}

	for _, v := range arr {
		if v == pivot {
			equal = append(equal, v)
		}
		if v < pivot {
			left = append(left, v)
		}
		if v > pivot {
			right = append(right, v)
		}
	}
	return append(append(QuickSort(left), equal...), QuickSort(right)...)
}

func QuickSortInPlace(arr []int) {
	if len(arr) <= 1 {
		return
	}

	pivot := arr[len(arr)/2]
	left, right := 0, len(arr)-1

	for left <= right {
		for arr[left] < pivot {
			left++
		}
		for arr[right] > pivot {
			right--
		}
		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}

	QuickSortInPlace(arr[:right+1])
	QuickSortInPlace(arr[left:])
}
