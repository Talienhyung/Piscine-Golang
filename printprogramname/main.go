package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[0]
	for i, r := range name {
		if i > 1 {
			z01.PrintRune(r)
		}
	}
	z01.PrintRune('\n')
}
