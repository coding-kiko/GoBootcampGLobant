package MyErrors

import (
	"fmt"
)

type Internal struct {
	description string
}

func (i *Internal) Error() string {
    return fmt.Sprintf("Error: ThirdParty: %s", i.description)
}

type ThirdParty struct {
	description string
}

func (t *ThirdParty) Error() string {
    return fmt.Sprintf("Error: ThirdParty: %s", t.description)
}

type Other struct {
	description string
}

func (o *Other) Error() string {
    return fmt.Sprintf("Error: ThirdParty: %s", o.description)
}

// function that accepts any type of error and prints its type
func CheckErrType(e error) {
	switch e.(type) {
	case *ThirdParty:
		fmt.Println("This error is a ThirdParty error")
	case *Internal:
		fmt.Println("This error is an Internal error")
	default:
		fmt.Println("This error is a different type of error (Other)")
	}
}
