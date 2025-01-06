package main

import "fmt"

func main() {
	p := 12
	q := 13

	// Arithmetic operations
	sum := p + q
	fmt.Println("Sum:", sum)

	mul := p * q
	fmt.Println("Multiplication:", mul)

	dev := p / q
	fmt.Println("Division (integer):", dev)

	mod := p % q
	fmt.Println("Modulo:", mod)

	// ^ is a bitwise XOR operator, not exponentiation
	bitwiseXor := p ^ q
	fmt.Println("Bitwise XOR:", bitwiseXor)

	// Loop from 0 to 5
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	// Call util() function from a.go
	util()
}

func util() {
	var day int
	fmt.Print("Enter the number between 1 and 7: ") // Fixed prompt
	fmt.Scan(&day)                                  // Take input from the user

	switch day {
	case 1:
		fmt.Println("Sunday")
	case 2:
		fmt.Println("Monday") // Corrected the sequence of days
	case 3:
		fmt.Println("Tuesday")
	case 4:
		fmt.Println("Wednesday")
	case 5:
		fmt.Println("Thursday")
	case 6:
		fmt.Println("Friday")
	case 7:
		fmt.Println("Saturday")
	default:
		fmt.Println("Invalid input. Please enter a number between 1 and 7.")
	}
}
