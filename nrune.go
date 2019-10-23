package piscine

func NRune(s string, n int) rune {
	runeArray := []rune(s)

	for index, value := range runeArray {
		if index == n-1 {
			return value
		}
	}
	return 0
}
