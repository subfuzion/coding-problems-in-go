package string_compression

import (
	"strconv"
	"strings"
)

func compress(s string) string {
	// use a string builder for efficiency while building compressed string
	b := strings.Builder{}

	// string as a character (rune) array
	chars := []rune(s)

	// same character counter
	n := 0

	for i := 0; i < len(chars); i++ {
		n++
		if i+1 >= len(chars) || chars[i] != chars[i+1] {
			b.WriteRune(rune(chars[i]))
			b.WriteString(strconv.Itoa(n))
			n = 0
		}
	}

	// don't return compressed string unless it's actually shorter than the original
	if len(s) <= b.Len() {
		return s
	}

	return b.String()
}
