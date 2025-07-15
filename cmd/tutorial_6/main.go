package main

import (
	"fmt"
)

func main() {
	var myRune = 'a'
	fmt.Printf("The rune is: %c\n", myRune)
	fmt.Printf("The rune as an integer is: %d\n", myRune)
	fmt.Printf("The rune as a hexadecimal is: %x\n", myRune)
	fmt.Printf("myRune is of type: %T\n", myRune)
	fmt.Printf("myRune = %v\n", myRune)

	var strSlice = []string{"u", "m", "e", "s", "h", " ", "k"}
	var catStr = ""
	for i := range strSlice {
		catStr += strSlice[i]
	}
	fmt.Printf("Concatenated string: %s\n", catStr)
}
