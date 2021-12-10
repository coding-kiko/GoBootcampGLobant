package main

import (
	"fmt"
	"strings"
)

/* YES, Go supports pointers! This allows you to pass values around different functions
by either value or by reference. This means that you are either going to be reasigning
the variable (by value) or passing the actual value that the piece of memory is 
pointing to (by referencing a memory address). An example: */

func main() { // set up the main function that will call two string modifying functions 
	s := "unchanged"
	new := "changed"
	ByValue(s, new) // I only pass the value of s
	fmt.Println(s) // output: unchanged

	ByReference(&s, new) // here I pass the memory address where the value is stored in memory 
	fmt.Println(s) // output: changed
}

func ByValue(s string, new string) {
	s = strings.Replace(s, s, new, 1) // inside this function s will have "changed" but wont change the value in 'main'
}

func ByReference(s *string, new string) {
	*s = strings.Replace(*s, *s, new, 1) // here I am dereferencing the pointer and reasigning the new value to where
										 // it points in memory, changing what the variable actually points to
}