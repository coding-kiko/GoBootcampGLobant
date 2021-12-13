package MyErrors

import (
	"errors"
	"fmt"
	"os"
)

type Internal struct {
	description string
}

func (i *Internal) Error() string {
    return fmt.Println("Error: ThirdParty:", i.description)
}

type ThirdParty struct {
	description string
}

func (t *ThirdParty) Error() string {
    return fmt.Println("Error: ThirdParty:", t.description)
}

type Other struct {
	description string
}

func (o *Other) Error() string {
    return fmt.Println("Error: ThirdParty:", o.description)
}

// function that accepts any type of error and prints its type
func CheckErrType(e error) {
	switch tp := e.(type) {
	case *ThirdParty:
		fmt.Println("This error is a ThirdParty error")
	case *Internal:
		fmt.Println("This error is an Internal error")
	default:
		fmt.Println("This error is a different type of error (Other)")
	}
}