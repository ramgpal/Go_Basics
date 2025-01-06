package main

import "fmt"

func main() {
	fmt.Println("Hii")

	// var ptr *int
	// fmt.Println("The value  of pointer is: ", ptr)

	myNumber := 23
	var ptr = &myNumber
	fmt.Println(ptr)
	// fmt.Println(&myNumber)

	fmt.Println(*ptr)
	// fmt.Println(*&myNumber)

	*ptr = *ptr * 2
	fmt.Println(myNumber)
}
