package datastructures

import (
	"fmt"
	"strings"
)

// Matrix of integers
type Matrix [][]int

// Equal tests that the other matrix has the same number of rows and columns
// containing the same values
func (m Matrix) Equal(other Matrix) bool {
	if other == nil {
		return false
	}

	if len(m) != len(other) || len(m[0]) != len(other[0]) {
		return false
	}

	for i, row := range m {
		for j, val := range row {
			if val != other[i][j] {
				return false
			}
		}
	}

	return true
}

// Square returns true if the number of rows and columns are equal
func (m Matrix) Square() bool {
	if len(m) != 0 && len(m) == len(m[0]) {
		return true
	}
	return false
}

// String returns the matrix as a formatted string
func (m Matrix) String() string {
	b := strings.Builder{}
	for _, row := range m {
		for _, val := range row {
			b.WriteString(fmt.Sprintf("%4d", val))
		}
		b.WriteString("\n")
	}
	return b.String()
}
