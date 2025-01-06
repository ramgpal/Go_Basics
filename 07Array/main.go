package main

import "fmt"

func main() {
	fmt.Println("Welcome to array chapter in Goland")

	var arr [4]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4

	fmt.Println(arr)

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
