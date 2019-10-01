package string_compression

import (
	"testing"

	"github.com/subfuzion/coding-problems-in-go/test"
)

// F is the type of function that implements the solution
type F func(string) string

// Test defines a test case
type Test struct {
	Input    string
	Expected string
}

// Tests is a test suite
var Tests = []Test{
	{
		"aabcccccaaa",
		"a2b1c5a3",
	},
	{
		"a",
		"a", // not "a1"
	},
	{
		"aa",
		"aa", // not "a2"
	},
	{
		"aaa",
		"a3",
	},
	{
		"hello world",
		"hello world", // not "h1e1l2o1 1w1o1r1l1d1"
	},
}

func TestSolution(t *testing.T) {
	for _, f := range Solutions {
		t.Run(test.GetFileFuncName(f), func(t *testing.T) {
			for _, test := range Tests {
				actual := compress(test.Input)
				if actual != test.Expected {
					t.Errorf("sample: '%s', expected: '%s', actual: '%s'", test.Input, test.Expected, actual)
				}
			}
		})
	}
}

// TODO - add your own solutions to the list below
// Solutions maps descriptive names to solution implementations
var Solutions = []F{
	compress,
	// compress1,
	// ...
}

//TODO - implement (and update the Solutions list, above)
//nolint:deadcode,unused //golangci-lint
/******************************************************************************
Compress string using the count of repeated characters. If compressed
string will not be smaller than the original, return the original.
Ex: "aabcccccaaa"
=>  "a2blc5a3"
Assume:
- the string consists only of uppercase and lowercase letters (a - z).
- the rune array is large enough to hold the extra characters after conversion
Hints (Cracking the Coding Interview 6): #92, #110
******************************************************************************/
func compress1(runes []rune, nRunes int) {
}
