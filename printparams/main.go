package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	for i := range arguments {
		for _, j := range arguments[i] {
			z01.PrintRune(j)
		}
		z01.PrintRune(10)
	}
}
