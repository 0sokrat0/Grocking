package quicksort_test

import (
	quicksort "go-blockchain/sort/QuickSort"
	"sort"
	"testing"
)

// Единица времени	Символ	Соотношение с секундой
// Наносекунда	ns	1 ns = 0.000000001 s
// Микросекунда	µs	1 µs = 0.000001 s
// Миллисекунда	ms	1 ms = 0.001 s
// Секунда	s	1 s = 1 s
// Минута	m	1 m = 60 s
// Час	h	1 h = 3600 s

func BenchmarkQuickSort(b *testing.B) {
	arr := []int{33, 2, 44, 1, 3, 2, 6, 7}
	for i := 0; i < b.N; i++ {
		quicksort.QuickSort(arr)
	}
}

func BenchmarkSortInts(b *testing.B) {
	arr := []int{33, 2, 44, 1, 3, 2, 6, 7}
	for i := 0; i < b.N; i++ {
		sort.Ints(arr)
	}
}

func BenchmarkQuickSortInPlace(b *testing.B) {
	arr := []int{33, 2, 44, 1, 3, 2, 6, 7}
	for i := 0; i < b.N; i++ {
		quicksort.QuickSortInPlace(arr)
	}
}

// func main() {

// 	arr := []int{33, 2, 44, 1, 3, 2, 6, 7}

// 	var m1 runtime.MemStats
// 	runtime.ReadMemStats(&m1)

// 	start := time.Now()

// 	sortedArr := quicksort.QuickSort(arr)

// 	var m2 runtime.MemStats
// 	runtime.ReadMemStats(&m2)

// 	memUsed := m2.Alloc - m1.Alloc

// 	elapsed := time.Since(start)

// 	fmt.Println("Отсортированный массив:", sortedArr)
// 	fmt.Println("Время выполнения:", elapsed)
// 	fmt.Println("Использовано памяти:", memUsed, "байт")
// }
