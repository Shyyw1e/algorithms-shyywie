package exercises

import (
	"fmt"
)

func Exersice1() {
	var target int
	var length int
	fmt.Println("Write down integer")
	fmt.Scan(&target)
	fmt.Println("Write down array length")
	fmt.Scan(&length)
	arr := make([]int, length)
	fmt.Println("Write array elements")
	for i := 0; i < length; i++ {
		var el int
		fmt.Scan(&el)
		arr[i] = el
	}
	res := -1
	for idx, elem := range arr {
		if elem == target {
			res = idx
			break
		}
	}
	fmt.Printf("Answer: %v\n", res)
}
