package piscine

func FirstRune(s string) rune {
	runeArray := []rune(s)
	return runeArray[len(s)-1]
}
