package main

import (
	"fmt"
)

func main() {
	// Take Input and Hello world program
	// var a int
	// fmt.Scanln(&a)
	// fmt.Println(a)
	// fmt.Println("Hello World")

	// var i int
	// for i = 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// i := 11
	// if i == 10 {
	// 	fmt.Println("If block is running")
	// } else {
	// 	fmt.Println("Else Block is running")
	// }

	i := 0
	for i < 10 {
		fmt.Println(i)
		i = i + 1
	}

	const a int = 20
	fmt.Println(a)
}
