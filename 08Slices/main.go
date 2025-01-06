package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hi, Now we will study about slices in Go")

	var slice = []string{"apple", "Banana", "Mango"}
	fmt.Println(slice)

	// to add value in the list
	slice = append(slice, "Papaya")
	fmt.Println(slice)

	slice = append(slice[1:3])
	fmt.Println(slice)

	tmpSlice := make([]int, 4)
	tmpSlice[0] = 234
	tmpSlice[1] = 150
	tmpSlice[2] = 275
	tmpSlice[3] = 390

	fmt.Println(tmpSlice)

	tmpSlice = append(tmpSlice, 555, 666, 723)
	fmt.Println(tmpSlice)

	sort.Ints(tmpSlice)

	fmt.Println(tmpSlice)

	// how to remove value from slices based on indexes
	courses := []string{"reactJS", "javascript", "swift", "python", "ruby"}
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
}
