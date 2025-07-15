package main

import "fmt"

type gasEngine struct {
	mpg uint8
	gallons uint8
}

func (e gasEngine) getRange() uint8 {
	return e.mpg * e.gallons
}

func canMakeIt(e gasEngine, distance uint8) bool {
	return e.getRange() >= distance
}

func main() {
	var myEngine gasEngine = gasEngine{22, 3}
	fmt.Println(myEngine.mpg, myEngine.gallons)
	fmt.Println("Range of myEngine:", myEngine.getRange())
	if canMakeIt(myEngine, 80) {
		fmt.Println("We can make it!")
	} else {
		fmt.Println("We cannot make it!")
	}


	var myEngine2 = struct {
		mpg uint8
		gallons uint8
	}{23, 24}	
	fmt.Println(myEngine2.mpg, myEngine2.gallons)
}