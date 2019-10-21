package piscine

func IteractiveFactorial(nb int) int {

	if nb >= 0 && nb <= 20 {
		number := 1
		for i := 1; i <= nb; i++ {
			number = number * i
		}
		return (number)
	} else {
		return 0
	}

}
