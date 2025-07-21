package main

import "fmt"

func main()  {
	var slice = []int32{1, 2, 3, 4, 5}
	var sliceCopy = slice
	sliceCopy[0] = 10
	fmt.Println("Original slice:", slice)
	fmt.Println("Modified copy:", sliceCopy)
	fmt.Println("Length of original slice:", len(slice))
	fmt.Println("Capacity of original slice:", cap(slice))
	fmt.Println("Length of copy slice:", len(sliceCopy))
	fmt.Println("Capacity of copy slice:", cap(sliceCopy))
	fmt.Println("Address of original slice:", &slice)
	fmt.Println("Address of copy slice:", &sliceCopy)	
}