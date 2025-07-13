package main

import "fmt"

func main() {
	var intArr [5]uint32
	fmt.Println("Length of the array:", len(intArr))
	fmt.Println("Capacity of the array:", cap(intArr))
	fmt.Println("Elements of the array:", intArr)
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:])

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])
	fmt.Println(&intArr[3])
	fmt.Println(&intArr[4])
}