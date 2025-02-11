package exercises

import(
	"fmt"
)

func Exersice2() {
	var	input string
	fmt.Println("Input string")
	fmt.Scan(&input)
	arr := []byte(input)
	l := len(arr)
	arr2 := make([]byte, l)
	j := 0
	for i := l - 1; i >= 0; i-- {
		arr2[i] = arr[j]
		j++
	}
	if string(arr) == string(arr2) {
		fmt.Println("Palyndrome")
	}else {
		fmt.Println("Not palyndrome")
	}
}