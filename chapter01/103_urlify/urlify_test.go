package url

import (
	"testing"
	"unicode/utf8"

	"github.com/subfuzion/coding-problems-in-go/test"
)

// F is the type of function that implements the solution
type F func(runes []rune, nRunes int)

// Test defines a test case
type Test struct {
	Input    string
	Expected string
}

// Tests is a test suite
var Tests = []Test{
	Test{"hello world", "hello%20world"},
	Test{"how now, brown cow", "how%20now,%20brown%20cow"},
}

func TestSolution(t *testing.T) {
	for _, f := range Solutions {
		t.Run(test.GetFileFuncName(f), func(t *testing.T) {
			for _, test := range Tests {
				// for this exercise we are supposed to work with character arrays
				// we can assume that we have a character array large enough to hold the
				// correct result; so we will use the actual answer to determine the size
				// needed, and then populate it with the input
				nRunes := utf8.RuneCountInString(test.Expected)
				runes := make([]rune, nRunes)
				for i, r := range test.Input {
					runes[i] = r
				}

				f(runes, utf8.RuneCountInString(test.Input))
				actual := string(runes)
				expected := test.Expected

				if actual != expected {
					t.Errorf("actual: '%s', expected: '%s'", actual, expected)
				}
			}
		})
	}
}

// TODO - add your own solutions to the list below
// Solutions maps descriptive names to solution implementations
var Solutions = []F{
	urlify,
	// urlify1,
	// ...
}

//TODO - implement (and update the Solutions list, above)
//nolint:deadcode,unused //golangci-lint
/******************************************************************************
/*
URLify: Write a method to replace all spaces in a string with '%20'.
You may assume that the string has sufficient space at the end to hold the
additional characters, and that you are given the "true" length of the string.
EXAMPLE
Input:  "Mr John Smith ", 13
Output: "Mr%20John%20Smith"
******************************************************************************/
func urlify1(runes []rune, nRunes int) {
}
