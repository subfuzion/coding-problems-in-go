package datastructures

import (
	"fmt"
	"strings"
)

type Matrix [][]int

func (m Matrix) Equal(m2 Matrix) bool {
	if m2 == nil {
		return false
	}

	if len(m) != len(m2) || len(m[0]) != len(m2[0]) {
		return false
	}

	for i, row := range m {
		for j, val := range row {
			if val != m2[i][j] {
				return false
			}
		}
	}

	return true
}

func (m Matrix) Square() bool {
	if len(m) != 0 && len(m) == len(m[0]) {
		return true
	}
	return false
}

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
