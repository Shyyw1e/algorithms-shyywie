package exercises

import (
	"fmt"
)

var fibResults = map[int]int{1:1, 2:1}

func Fibonacci(num int) int{
	if num <= 0 {
		return 0 // Защита от некорректного ввода
	}
	if val, ok := fibResults[num]; ok {
		return val // Если значение уже есть, просто возвращаем
	}
	fibResults[num] = Fibonacci(num-1) + Fibonacci(num-2)
	return fibResults[num]
	
}

func Exersice3() {
	var num int
	fmt.Println("Enter Fibonacci index")
	fmt.Scan(&num)
	
	fmt.Printf("Answer: %v\n", Fibonacci(num))
	
	
}