package main

import "fmt"

func main()  {
	var p *int32 = new(int32)
	var i int32
	fmt.Printf("The value p points to is: %v", *p)
	fmt.Printf("\nThe value of i is: %v", i)
	p = &i
	*p = 42
	fmt.Printf("\nThe value p points to is now: %v", *p)
	fmt.Printf("\nThe value of i is now: %v", i)
	
}