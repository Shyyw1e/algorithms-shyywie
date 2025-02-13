package exercises

import "fmt"


func SetArray(arr []int) []int {
	seen := make(map[int]bool)
	var result []int

	for _, num := range arr {
		if !seen[num] {
			seen[num] = true
			result = append(result, num)
		}
	}
	return result
}

func Exersice6() {
	var length int
	fmt.Println("Enter array length")
	fmt.Scan(&length)

	arr := make([]int, length)
	fmt.Println("Enter array elements")
	
	for i := 0; i < length; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println(SetArray(arr))

}