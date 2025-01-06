package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Println("Hello,", name)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza: ")

	// comma ok/error ok
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of rating is %T", input)
}
