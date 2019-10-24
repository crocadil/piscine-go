package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	for index, arg := range args {
		if index == 0 {
			continue
		}
		for _, c := range arg {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
	}
}
