//Filename:main.go
//Waitgroups:Another mechanism of concurrency and synchronization
//Keeps track of goroutines and makes main() to wait until they are executed
//Does not require time.sleep()

package main

import (
	"fmt"
	"sync"
)

func festival(wg *sync.WaitGroup) {

	fmt.Println("Christmas is the best annual christian festival")
	wg.Done() //decrementing by 1
}

func wish(wg *sync.WaitGroup) {

	fmt.Println("I wish a had a mansion")
	wg.Done() //decrementing by 1
}

// main is a goroutine
func main() {
	//create a waitgroup
	var wg sync.WaitGroup
	//increment waitgroup by 1
	wg.Add(1)

	//calls goroutine festival()
	go festival(&wg)

	//increment waitgroup 1
	wg.Add(1)

	//calls goroutine wish()
	go wish(&wg)

	//makes main()to wait for wg to become 0
	wg.Wait()

	fmt.Println("Happy Birthday to you!")

}
