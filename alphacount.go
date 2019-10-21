package piscine

func AlphaCount(str string) int {
	counter := 0

	for _, i := range str {
		if i >= 'a' && i <= 'z' || i >= (65) && i <= (90) {
		counter++
		}
	}
	return counter

}
