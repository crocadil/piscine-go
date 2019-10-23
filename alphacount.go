package piscine

func AlphaCount(str string) int {
	k := 0
	for _, value := range str {
		if value >= 'a' && value <= 'z' || value >= 'A' && value <= 'Z' {
			k++
		}
	}
	return k
}
