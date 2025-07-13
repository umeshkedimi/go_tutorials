package main

import "fmt"

func main() {
	var intArr [4]int32 = [4]int32{1,2}
	fmt.Println(intArr)
	strArr := [4]string{"code", "is", "poetry", "!"}
	fmt.Printf("%sing %s %s%s\n", strArr[0], strArr[1], strArr[2], strArr[3])

	intSlice := []int32{1, 2, 7}
	fmt.Printf("intSlice: %v, len: %d, cap: %d\n", intSlice, len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 34, 2, 45, 76, 89)
	fmt.Printf("intSlice: %v, len: %d, cap: %d\n", intSlice, len(intSlice), cap(intSlice))
	fmt.Println(intSlice)

	var intSlice2 []int32 = []int32{8,9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	var intSlice3 []int32 = make([]int32, 5, 12)
	fmt.Printf("intSlice3: %v, len: %d, cap: %d\n", intSlice3, len(intSlice3), cap(intSlice3))
}