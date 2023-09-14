package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if args[0] == "" {
		return
	}
	upper := false
	if len(args) > 0 && args[0] == "--upper" {
		upper = true
		args = args[1:]
	}
	for _, arg := range args {
		if isValidNumber(arg) {
			num := atoi(arg)
			var letter rune
			if num >= 1 && num <= 26 {
				letter = rune('a' + num - 1)
				if upper {
					letter = rune('A' + num - 1)
				}
				z01.PrintRune(letter)
			} else {
				z01.PrintRune(' ')
			}
		} else {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}

func isValidNumber(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func atoi(s string) int {
	result := 0
	for _, c := range s {
		result = result*10 + int(c-'0')
	}
	return result
}
