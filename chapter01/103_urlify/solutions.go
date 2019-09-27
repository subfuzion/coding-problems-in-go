package url

func urlify(runes []rune, nRunes int) {
	ri := len(runes) - 1
	for i := nRunes - 1; i >= 0; i-- {
		if runes[i] == ' ' {
			runes[ri-2] = '%'
			runes[ri-1] = '2'
			runes[ri] = '0'
			ri -= 3
		} else {
			runes[ri] = runes[i]
			ri--
		}
	}
}
