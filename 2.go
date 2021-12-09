/* defer is a keyword used to define an function call that you may want to execute after the
surrounding function finishes executing every single piece of code. Under the hood, what is
happening is each function called with defer, is being pushed to a call stack, which then
will execute in a reversed order (The first introduced will be the last executed) */

// lets say we want to open an x amount of files, from inside a function
func main() {
	a, err1 := os.Open(file1)
	if err1 != nil {
		return
	}
	defer a.Close()

	b, err2 := os.Open(file2)
	if err2 != nil {
		return
	}
	defer b.Close()

	c, err3 := os.Open(file3)
	if err3 != nil {
		return
	}
	defer c.Close()

	// Here we would perform any operation with the content of the files but without any worries
	// because we defered the closing of them, meaning that after this piece of code is executed
	// and the main function returns, all of the files will be closed.

	return x
}