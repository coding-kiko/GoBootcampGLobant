package main

import (
	"fmt"
	MyErrors "github.com/coding-kiko/GoBootcampGlobant/12-errorPackage/MyErrors"
)

func main() {
	// creating a 'third party error, that could be trhown by a function'
	TPerr := &MyErrors.ThirdParty{Description: "Table name not found"}
	fmt.Println(TPerr) // output: Error: ThirdParty: Table name not found

	// same here but internal
	Ierr := &MyErrors.Internal{Description: "Unexpected '('"}
	fmt.Println(Ierr) // output: Error: Internal: Unexpected '('

	// 'Other' type of error
	Other := &MyErrors.Other{Description: "Other type of error"}
	fmt.Println(Other) // output: Error: Other type of error

	// testing the error checker function:
	MyErrors.CheckErrType(Ierr) // output: This is an Internal error
}