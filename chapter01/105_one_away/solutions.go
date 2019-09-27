package one_away

func isOneEditAway(a, b string) bool {
	if a == b {
		return true
	}

	d := len(a) - len(b)
	switch d {
	case 0:
		return oneEditReplace(a, b)
	case -1:
		return oneEditInsert(a, b)
	case 1:
		return oneEditInsert(b, a)
	default:
		// difference is too great for a single edit
		return false
	}
}

func oneEditReplace(a, b string) bool {
	found := false
	ra := []rune(a)
	rb := []rune(b)

	for i := range ra {
		if ra[i] != rb[i] {
			// if found a mismatch already, then this is more than one edit away
			if found {
				return false
			}
			found = true
		}
	}
	return found
}

func oneEditInsert(a, b string) bool {
	ra := []rune(a)
	rb := []rune(b)

	for ia, ib := 0, 0; ib < len(ra); {
		if ra[ia] == rb[ib] {
			ia++
			ib++
		} else {
			ib++
		}

		if ib-ia > 1 {
			return false
		}
	}

	return true
}
