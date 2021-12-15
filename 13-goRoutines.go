package main

/* Two or more functions can be made concurrent in go using goroutines. These can be seen as lightweight
threads which allow functions to coexist. Concurrency which is not to be confused with paralelism,
run independently on each other, that is except, we use some type of syncronization, which I will
discuss in task 14. The main function is the prime example of a goroutine, which starts at runtime.
Lets see an example: */

import (
	"fmt"
	"time"
)

func otherRoutine() {
	for i := 0; i < 10; i++ {
		time.Sleep(5 * time.Second) // prints every 5 seconds but concurrently with the main routine
		fmt.Print(" 5 seconds ")
	}
}

func main() {
	go otherRoutine() // I imediately start a second goroutine after starting main
	for i := 0; i < 10; i++ {
		time.Sleep(3 * time.Second) // every 3 seconds this loop will print 'Main'
		// if I dont wait here and let main finish, the otherRoutine
		// will not be waited and thge program will end
		fmt.Print("Main  ")
	}
}

// output: Main    5 seconds Main  Main    5 seconds Main    5 seconds Main
