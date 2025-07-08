package main

import "fmt"

func main() {
	var intNum int = 32767
	intNum = intNum + 1 // This will cause an overflow
	fmt.Println("The value of intNum is:", intNum)
	fmt.Println("The maximum value of int is:", int(^uint(0)>>1)) //

	var floatNum float64 = 12345678.9
	fmt.Println("The value of floatNum is:", floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 * float32(intNum32)
	fmt.Println("The result of floatNum32 * intNum32 is:", result)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1/ intNum2) // This will perform integer division
	fmt.Println(intNum1% intNum2) // This will perform modulus operation
}