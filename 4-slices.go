/* Arrays in Go are fixed sized structures that store a sequence of elements of the same type, Slices,
on the other hand, are dynamically allocated, meaning that you can 'extend' them at any time by
using the builtin function 'append()' */

func arrays() {
	var array [2]int // declaring array of size 2 (integers)
	array[0] = 3 // set a new value for the first index
	fmt.Println(array)	// output: [3 0]
}

func slices() {
	slice := make([]int, 2) // declaring slice of initial size 2
	slice = append(slice, 1) // appending new value to the end
	fmt.Println(slice) // output: [0 0 1]

	new := []string{"Juan", "Carlos", "Marcelo"} // declaring and initializing
	new = append(new, "Fernando")
	fmt.Println(new) // output: [Juan Carlos Marcelo Fernando]
}