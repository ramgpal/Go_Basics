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

	fmt.Print("Please provide your feedback rating: ")
	rating, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Thanks for the rating,", rating)
}
