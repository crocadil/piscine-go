package piscine

func Fibonacci(index int) int {

	if index <= 0 {
		return -1
	}
	return Fibonacci(index-1) + Fibonacci(index-2)

}
