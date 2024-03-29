package matrix_ops

import (
	"testing"

	"github.com/subfuzion/coding-problems-in-go/datastructures"

	"github.com/subfuzion/coding-problems-in-go/test"
)

// F is the type of function that implements the solution
type F func(datastructures.Matrix) bool

// Test defines a test case
type Test struct {
	Input    datastructures.Matrix
	Expected datastructures.Matrix
}

var Tests = []Test{
	{
		Input: datastructures.Matrix{
			{1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
		},
		Expected: datastructures.Matrix{
			{1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
		},
	},
	{
		Input: datastructures.Matrix{
			{1, 1, 1, 1, 1, 1},
			{1, 1, 0, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 0, 1},
			{1, 1, 1, 1, 1, 1},
		},
		Expected: datastructures.Matrix{
			{1, 1, 0, 1, 0, 1},
			{0, 0, 0, 0, 0, 0},
			{1, 1, 0, 1, 0, 1},
			{0, 0, 0, 0, 0, 0},
			{1, 1, 0, 1, 0, 1},
		},
	},
	{
		Input: datastructures.Matrix{
			{1, 0, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
			{0, 1, 1, 1, 1, 1},
		},
		Expected: datastructures.Matrix{
			{0, 0, 0, 0, 0, 0},
			{0, 0, 1, 1, 1, 1},
			{0, 0, 1, 1, 1, 1},
			{0, 0, 1, 1, 1, 1},
			{0, 0, 0, 0, 0, 0},
		},
	},
	{
		Input: datastructures.Matrix{
			{0, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
		},
		Expected: datastructures.Matrix{
			{0, 0, 0, 0, 0, 0},
			{0, 1, 1, 1, 1, 1},
			{0, 1, 1, 1, 1, 1},
			{0, 1, 1, 1, 1, 1},
			{0, 1, 1, 1, 1, 1},
		},
	},
}

func TestSolution(t *testing.T) {
	for _, f := range Solutions {
		t.Run(test.GetFileFuncName(f), func(t *testing.T) {
			for _, test := range Tests {
				actual := test.Input
				expected := test.Expected

				// if valid input then matrix is zerod in place
				ok := f(test.Input)

				if !ok {
					t.Errorf("error rotating matrix (must be square with non-zero dimensions):\n%v", test.Input)
				}

				if !actual.Equal(expected) {
					t.Errorf("actual != expected\nActual:\n%sExpected:\n%s", actual, expected)
				}
			}
		})
	}
}

// TODO - add your own solutions to the list below
// Solutions maps descriptive names to solution implementations
var Solutions = []F{
	nullify,
	// nullify1,
	// ...
}

//TODO - implement (and update the Solutions list, above)
//nolint:deadcode,unused //golangci-lint
/******************************************************************************
Zero Matrix: Write an algorithm such that if an element in an MxN matrix is 0,
its entire row and column are set to 0.
Hints (Cracking the Coding Interview 6): #17, #74, #702
******************************************************************************/
func nullify1(m datastructures.Matrix) bool {
	return false
}
