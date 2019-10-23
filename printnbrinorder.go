
package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {

	a := 0
	adiko := []rune{}
	z := 0

	if n == 0 {
		adiko = append	(adiko, rune(n))
	}

	for n != 0 {
		a = n % 10
		n = n / 10
		adiko = append	(adiko, rune(a))
	}

	for range	adiko {
		z++
	}

	for i := 0; i < z-1; i++ {
		for j := 0; j < z-1-i; j++ {
			if	adiko[j] >	adiko[j+1] {
				k :=	adiko[j]
				adiko[j] =	adiko[j+1]
				adiko[j+1] = k
			}
		}
	}

	for i := 0; i < z; i++ {
		z01.PrintRune	(adiko[i] + 48)
	}
	
}
