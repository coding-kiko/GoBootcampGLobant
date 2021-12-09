package main

import "fmt"

/* Unlike other languages, Go does not come with an integrated while loop, instead, the for loop can be used
without any new variable initialization or incrementing or decrementing. By that I mean, using only the
conditional statement, making whatever is happening inside the for loop to be controlled by the named
condition. Lets see an actual example: */

func main() {
	x := 36
	sqrt := 1

	// In this example I look for the square root of a squared number,
	// starting from 1 and going on until I find it, incrementing each iteration.
	for sqrt <= x {
		if sqrt*sqrt == x {
			fmt.Println(sqrt)
		}
		sqrt++
	}
}
