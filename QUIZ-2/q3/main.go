//Filename:main.go
//Demonstrate goroutines:functions that execute independently from others in a program

package main

import (
	"fmt"
	"time"
)

func festival() {
	fmt.Println("Christmas is the best annual christian festival")
}

func wish() {
	fmt.Println("I wish a had a mansion")
}

// main is a goroutine
func main() {

	go festival()
	go wish()

	//A temporary solution for challenges of concurrency
	time.Sleep(time.Second) //makes the execution of goroutines non-deterministic

	fmt.Println("Happy Birthday to you!")

}
