// Filename: main.go
// Demonstrate buffered channel:it has a capacity/buffer/memory space for storage unlike unbuffered
// Hence no blocking/deadlock because of the capability of storing values written into it
// However blocking/deadlock can occur by exceeding its capacity or not writing anything into it
// Note:no synchronization between writing and reading, as they canbe done at a later time
package main

import (
	"fmt"
)

func main() {
	//creates a channel
	decade := make(chan int, 4)

	//write data into the channel
	decade <- 2
	decade <- 4
	decade <- 6
	decade <- 8

	//loop through the values
	for i := 0; i < 4; i++ {

		//main() goroutine prints the data received from channel 'decade'
		fmt.Println(<-decade)
	}

}
