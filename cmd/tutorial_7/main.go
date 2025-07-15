package main

import "fmt"

type gasEngine struct {
	mpg uint8
	gallons uint8
}

func (e gasEngine) getRange() uint8 {
	return e.mpg * e.gallons
}

type electricEngine struct {
	kwh uint8
	charge uint8
}

func (e electricEngine) getRange() uint8 {
	return e.kwh * e.charge
}

type engine interface {
	getRange() uint8
}


func canMakeIt(e engine, miles uint8) {
	if miles <= e.getRange() {
		fmt.Println("We can make it!")
	} else {
		fmt.Println("We cannot make it!")
	}
}

func main() {
	var myEngine gasEngine = gasEngine{22, 5}
	fmt.Println(myEngine.mpg, myEngine.gallons)
	fmt.Println("Range of myEngine:", myEngine.getRange())
	canMakeIt(myEngine, 60)

	var myEngine3 = electricEngine{22, 4}
	fmt.Println(myEngine3.kwh, myEngine3.charge)
	fmt.Println("Range of myEngine3:", myEngine3.getRange())
	canMakeIt(myEngine3, 60)
}