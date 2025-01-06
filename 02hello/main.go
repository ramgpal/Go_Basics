package main

import "fmt"

func main() {
	var username string = "Ram"
	// short declaration -> only valid for local scope
	// username := "Ram"

	fmt.Println(username)

	fmt.Println("Variable is of Type: %T \n", username)
}
