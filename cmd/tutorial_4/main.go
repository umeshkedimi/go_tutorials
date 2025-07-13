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

	// map
	var myMap map[string]uint8 = make(map[string]uint8)
	myMap["umesh"] = 12
	myMap["suresh"] = 34
	myMap["rajesh"] = 56
	fmt.Println(myMap)
	fmt.Println(myMap["umesh"])
	fmt.Println(myMap["unknown"]) // returns 0 for non-existing key

	var myMap2 = map[string]uint8{"city": 123, "country": 255}
	fmt.Println(myMap2)

	var myMap3 = map[uint8]string{1: "one", 25: "two", 3: "three"}
	fmt.Println(myMap3)

	// delete(myMap2, "city")

	var city, ok = myMap2["city"]
	if ok {
		fmt.Println("City found:", city)
	} else {
		fmt.Println("City not found")
	}

	for key, value := range myMap {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	for name, val := range myMap2 {
		fmt.Printf("Name: %s Val: %d\n", name, val)
	}
}