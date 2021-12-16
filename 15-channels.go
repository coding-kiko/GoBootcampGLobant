package main

import (
	"fmt"
	conv "strconv"
	"time"
)

var flag bool = true // even though it is not the best practice I used a global variable to control
// the for loops inside each go routine

func sender(msg chan string) {
	for i := 1; flag; i++ {
		msg <- "Important info n° " + conv.Itoa(i) // Every second this goroutine sends a new message
		time.Sleep(time.Second)
	}
	fmt.Println("Main routine closing, returning sender.")
	return
}

func receiver(msg chan string) {
	for i := 1; flag; i++ {
		str := <-msg // every second this one grabas the message and unblocks the channel
		//leaving space for the next message
		fmt.Println(str)
		time.Sleep(time.Second)
	}
	fmt.Println("Main routine closing, returning receiver.")
	return
}

func main() {
	msg := make(chan string) // creating unbuffered chanel
	// starting the go routines
	go receiver(msg)
	go sender(msg)
	time.Sleep(5 * time.Second) // after 5 seconds, 5 messages were interchanged through msg
	flag = !flag                // invert logic to close for loops and return routines
	time.Sleep(time.Second)     // leave some time for the print statement
}

// output:
// Important info n° 1
// Important info n° 2
// Important info n° 3
// Important info n° 4
// Important info n° 5
// Main routine closing, returning receiver.
// Main routine closing, returning sender.
