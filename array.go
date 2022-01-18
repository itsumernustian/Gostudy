package main

import "fmt"

func main() {

	var arr [11]int
	for i := 0; i < 11; i++ {
		arr[i] = i * 2
	}
	fmt.Println(arr)

	arr1 := [11]int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	fmt.Println(arr1)
	arr2 := [...]int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22}
	fmt.Println(arr2)
}
