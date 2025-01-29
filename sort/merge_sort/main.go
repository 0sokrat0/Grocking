package main

import (
	"fmt"
)

func main() {

	// arr := []int{33, 2, 44, 1, 3, 2, 6, 7}

	table := []int{2, 3, 7, 8, 10}

	// for _, v := range arr {

	// 	fmt.Print(v, " ")
	// }

	// fmt.Print("\n")

	// for _, v := range arr {
	// 	time.Sleep(time.Second * 1)
	// 	fmt.Print(v, " ")
	// }

	// fmt.Print("\n")

	// for _, v := range arr {
	// 	fmt.Print(v*2, " ")
	// }

	// fmt.Print("\n")

	// for _, v := range arr {
	// 	if v == arr[0] {
	// 		time.Sleep(time.Millisecond * 1)
	// 		fmt.Print(v*2, " ")
	// 	}
	// }

	fmt.Println(tableMultiplication(table, 0))
}

func tableMultiplication(table []int, mult int) []int {

	if mult == len(table) {
		return table
	}

	for i := 0; i < len(table)-1; i++ {
		table[i] *= table[mult]
		print(table[i], " ")
	}
	return tableMultiplication(table, mult+1)
}
