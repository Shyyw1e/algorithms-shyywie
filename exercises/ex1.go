package exercises

import(
	"fmt"
)

func Exersice1(){
	var target int
	var length int
	fmt.Println("Write down integer")
	fmt.Scan(&target)
	fmt.Println("Write down array length")
	fmt.Scan(&length)
	arr := make([]int, length)

	fmt.Print(arr)
}