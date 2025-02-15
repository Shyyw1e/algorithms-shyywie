package exercises

import (
	"fmt"
)

var FibResults = map[int]int {1:1, 2:1}

func Fibonacci(num int) int {
	if num <= 0 {
		return 0
	}
	if val, ok := FibResults[num]; ok {
		return val
	}
	FibResults[num] = Fibonacci(num - 1) + Fibonacci(num - 2)
	return FibResults[num]
}

func Exersice3(){
	var num int
	fmt.Println("Enter Fibonacci number")
	fmt.Scan(&num)

	fmt.Printf("Answer: %v\n", Fibonacci(num))
}