package main

import("fmt"
	"errors")

func main() {
	var printValue string = "Hello, Go!"
	printMe(printValue)

	numerator := 10
	denominator := 5
	quotient, remainder, err := intDivision(numerator, denominator)
	
	switch {
	case err != nil:
		fmt.Printf(err.Error())
	case remainder == 0:
		fmt.Printf("The result of the integer division is %v\n", quotient)
	default:
		fmt.Printf("The result of the integer division is %v with a remainder of %v\n", quotient, remainder)
	}
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot divide by zero")
		return 0, 0, err // Avoid division by zero
	}
	quotient := numerator / denominator
	remainder := numerator % denominator
	return quotient, remainder, err
}