package main

import (
	"fmt"
	operations "github.com/coding-kiko/GoBootcampGlobant/8-calculator/operations" // importing the operations that are one folder ahead
	"os"
	"strconv"
)

// reads from cmd line and makes operation depening on the operatoin sign
// only +, -, x and / are implemented. example: ./calc 1 + 3
func main() {
	a, _ := strconv.Atoi(os.Args[1])
	op := os.Args[2]
	b, _ := strconv.Atoi(os.Args[3])
	switch op {
	case "+":
		fmt.Println(operations.Add(a, b))
	case "-":
		fmt.Println(operations.Subtract(a, b))
	case "/":
		fmt.Println(operations.Divide(a, b))
	case "x":
		fmt.Println(operations.Multiply(a, b))
	default:
		fmt.Println("Invalid operator")
	}
}