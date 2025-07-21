package main

import (
	"fmt"
	"time"
	"sync"
)

var m = sync.Mutex{}
var wg = sync.WaitGroup{}
var dbData =[]string{"data1", "data2", "data3", "data4", "data5"}
var results = []string{}

func main() {
	t0 := time.Now()
	for i:=0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("Total time taken: %v\n", time.Since(t0))
	fmt.Println("Results:", results)
}

func dbCall(i int) {
	// Simulate a database call
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is:", dbData[i])
	m.Lock()
	results = append(results, dbData[i])
	m.Unlock()
	wg.Done()
}