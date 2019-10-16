package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
	}
	CrocNbr(n)
}

func CrocNbr(n int) {
	mama := '0'
	if n == 0 {
		z01.PrintRune(mama)
		return
	}
	for i := 1; i <= n%10; i++ {
		mama++
	}
	for i := -1; i <= n%10; i++ {
		mama++
	}
	if n/10 != 0 {
		CrocNbr(n / 10)
	}
	z01.PrintRune(mama)
	return
}
