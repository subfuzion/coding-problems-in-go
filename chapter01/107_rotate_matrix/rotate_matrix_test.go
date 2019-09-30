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
			{0, 1, 2},
			{3, 4, 5},
			{6, 7, 8},
		},
		Expected: datastructures.Matrix{
			{6, 3, 0},
			{7, 4, 1},
			{8, 5, 2},
		},
	},
	{
		Input: datastructures.Matrix{
			{0, 1, 2, 3, 4, 5},
			{6, 7, 8, 9, 10, 11},
			{12, 13, 14, 15, 16, 17},
			{18, 19, 20, 21, 22, 23},
			{24, 25, 26, 27, 28, 29},
			{30, 31, 32, 33, 34, 35},
		},
		Expected: datastructures.Matrix{
			{30, 24, 18, 12, 6, 0},
			{31, 25, 19, 13, 7, 1},
			{32, 26, 20, 14, 8, 2},
			{33, 27, 21, 15, 9, 3},
			{34, 28, 22, 16, 10, 4},
			{35, 29, 23, 17, 11, 5},
		},
	},
}

func TestSolution(t *testing.T) {
	for _, f := range Solutions {
		t.Run(test.GetFileFuncName(f), func(t *testing.T) {
			for _, test := range Tests {
				actual := test.Input
				expected := test.Expected

				// if valid input then matrix is rotated in place
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
	rotate,
	// rotate1,
	// ...
}

//TODO - implement (and update the Solutions list, above)
//nolint:deadcode,unused //golangci-lint
/******************************************************************************
Rotate Matrix: Given an image represented by an NxN matrix, where each pixel in
the image is 4 bytes, write a method to rotate the image by 90 degrees. Can you
do this in place?
******************************************************************************/
func rotate1(m datastructures.Matrix) bool {
	return false
}
