package string_rotation

func isStringRotation(a, b string) bool {
	n := len(a)
	if n == 0 || n != len(b) {
		return false
	}

	aa := a + a
	// check that b is a substring of aa (as will always be the case if b is a valid rotation)
	return isSubstring(aa, b)
}

func isSubstring(a, b string) bool {
	return index(a, b) >= 0
}

// see https://golang.org/src/strings/strings_amd64.go
// for industrial strength solution
func index(a, b string) int {
	n := len(b)
	switch {
	case n == 0:
		return 0
	case n == len(a):
		if b == a {
			return 0
		}
		return -1
	case n > len(a):
		return -1
	}

	ra := []rune(a)
	rb := []rune(b)

	for i := 0; i < len(ra)-len(rb); i++ {
		for j := 0; j < len(rb); j++ {
			if rb[j] != ra[i+j] {
				break
			}
			if j+1 == len(rb) {
				return i
			}
		}
	}

	return -1
}
