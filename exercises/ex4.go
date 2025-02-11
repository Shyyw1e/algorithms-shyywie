package exercises

import (
	"fmt"
	"math"
)

func Exersice4() {
	var length int
	fmt.Println("Enter array length")
	fmt.Scan(&length)
	if length < 2 {
		fmt.Println("Array must have at least 2 elements")
		fmt.Scan(&length)
		return
	}
	arr := make([]int, length)
	fmt.Println("Enter array elements")

	for i := 0; i < length; i++ {
		var el int
		fmt.Scan(&el)
		arr[i] = el
	}
	maxNum := math.MinInt
	maxNum2 := math.MinInt
	for _, el := range arr {
		if el > maxNum {
			maxNum2 = maxNum
			maxNum = el
		}else if el > maxNum2 && el < maxNum{
			maxNum2 = el
		}
	}
	fmt.Printf("Answer: %v\n", maxNum2)
}