package piscine

func LastRune(s string) rune {
	runeArray := []rune(s)
	count := 0
	for s := range runeArray {
		count = s + 1
	}
	return runeArray[count-1]
}
