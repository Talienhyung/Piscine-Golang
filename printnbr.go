package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
	}
	printDigits(n)
}

func printDigits(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	if n/10 != 0 {
		printDigits(n / 10)
		z01.PrintRune(rune(n%10 + '0'))
	}
}
