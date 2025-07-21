package main

import (
	"fmt"
	"time"
	"sync"
)


var wg = sync.WaitGroup{}
var dbData =[]string{"data1", "data2", "data3", "data4", "data5"}


func main() {
	t0 := time.Now()
	for i:=0; i < 1000; i++ {
		wg.Add(1)
		go count()
	}
	wg.Wait()
	fmt.Printf("Total time taken: %v\n", time.Since(t0))
	// fmt.Println("Results:", results)
}

func count() {
	var res int
	for i:= 0; i < 100000000; i++ {
		res += i
	}
	wg.Done()
}

