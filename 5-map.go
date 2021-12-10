package main

import "fmt"

/* We can search for a specific key ina map, and it will return two values: the value of the key
and a boolean indicating if it exists: */ 

func main() {
	person := make(map[string]int) // create map
	person["age"] = 56 // add new key - value pairs
	person["id"] = 583512

	val, exists := person["age"]
	fmt.Printf("Age: %d - exists: [%t]\n", val, exists) // output: Age: 56 - exists: [true]

	_, exists = person["nil"] // searching for a key thata doesn't exist
	fmt.Printf("exists: [%t]\n", exists) // output: exists: [false]
}