package main

import "fmt"

func main()  {
	var thing1 = [5]float64{1.0, 2.0, 3.0, 4.0, 5.0}
	fmt.Printf("The memory location of the thing1 array is: %p", &thing1)
	var result [5]float64 = square(thing1)
	fmt.Printf("\nThe result of squaring the array is: %v", result)
	fmt.Printf("\nThe memory location of the result array is: %p", &result)
	fmt.Printf("\nThe memory location of the thing1 array after squaring is: %p", &thing1)
}

func square(thing2 [5]float64) [5]float64 {
	fmt.Printf("\nThe memory location of the thing2 array is: %p", &thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return thing2
}