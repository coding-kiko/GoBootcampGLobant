package MyErrors

import (
	"fmt"
)

// new error type with a single description
type Internal struct {
	Description string
}
// implementing Error() method: now my new type is considered as an error (thanks interface gods)
func (i *Internal) Error() string {
    return fmt.Sprintf("Error: Internal: %s", i.Description)
}

// new error type with a single description
type ThirdParty struct {
	Description string
}

func (t *ThirdParty) Error() string {
    return fmt.Sprintf("Error: ThirdParty: %s", t.Description)
}

// new error type with a single description
type Other struct {
	Description string
}

func (o *Other) Error() string {
    return fmt.Sprintf("Error: %s", o.Description)
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
