package MyErrors

import (
	"fmt"
)

type Internal struct {
	Description string
}

func (i *Internal) Error() string {
    return fmt.Sprintf("Error: ThirdParty: %s", i.Description)
}

type ThirdParty struct {
	Description string
}

func (t *ThirdParty) Error() string {
    return fmt.Sprintf("Error: ThirdParty: %s", t.Description)
}

type Other struct {
	Description string
}

func (o *Other) Error() string {
    return fmt.Sprintf("Error: ThirdParty: %s", o.Description)
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
