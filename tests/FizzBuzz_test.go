package tests

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func FizzBuzz(n int) string {
	switch {
	case n%5 == 0 && n%3 == 0:
		return "FizzBuzz"
	case n%5 == 0:
		return "Buzz"
	case n%3 == 0:
		return "Fizz"
	default:
		return strconv.Itoa(n)
	}
}

func TestFizzBuzz(t *testing.T) {
	testcases := []struct { // create test cases struct
		name     string // string description
		input    int    // function input
		expected string // expected output
	}{ // anonymous struct instances for testing
		{name: "validate multiple of 3", input: 9, expected: "Fizz"},
		{name: "validate multiple of 5", input: 20, expected: "Buzz"},
		{name: "validate multiple of 3 and 5", input: 15, expected: "FizzBuzz"},
		{name: "validate multiple of 3", input: 3, expected: "Fizz"},
		{name: "validate multiple of 5", input: 25, expected: "Buzz"},
		{name: "validate multiple of 3 and 5", input: 60, expected: "FizzBuzz"},
		{name: "validate multiple of none", input: 8, expected: "8"},
		{name: "validate multiple of none", input: 88, expected: "88"},
	}
	for _, tCase := range testcases { // iterate over every wantes test case
		t.Run(tCase.name, func(t *testing.T) {
			result := FizzBuzz(tCase.input)
			assert.Equal(t, tCase.expected, result) // check if result is the same as expected
		})
	}
}
