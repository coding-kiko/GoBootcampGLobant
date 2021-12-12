package main

import (
	"fmt"
	"reflect"
)

type fake struct{}

func main() {
	isPrimitive("this string")
	isPrimitive(0.84)
	isPrimitive(42)
	isPrimitive(true)
	isPrimitive(fake{})
	isPrimitive(nil)
}

// by definition an empty interface will be implemented by any type, so by doing this we can accept
// any type of variable into the function
func isPrimitive(x interface{}) {
	switch tp := x.(type) { // here we use go's type switches, which enables us to make multiple type assertions
	case string:
		fmt.Println(x, "is of primitive type:", reflect.TypeOf(tp))
	case int:
		fmt.Println(x, "is of primitive type:", reflect.TypeOf(tp))
	case float64:
		fmt.Println(x, "is of primitive type:", reflect.TypeOf(tp))
	case bool:
		fmt.Println(x, "is of primitive type:", reflect.TypeOf(tp))
	default:
		fmt.Println(x, "is not of primitive type:", reflect.TypeOf(tp))
	}
	return 
}