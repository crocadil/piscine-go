package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
	}
	Adika(n)
}

func Adika(n int) {
	lon := '0'
	if n == 0 {
		z01.PrintRune(lon)
		return
	}
	for i := 1; i <= n%10; i++ {
		lon++
	}
	for i := -1; i <= n%10; i++ {
		lon++
	}
	if n/10 != 0 {
		Inks(n / 10)
	}
	z01.PrintRune(lon)
	return
}

func main() {
	PrintNbr(-123)
	PrintNbr(0)
	PrintNbr(123)
}
