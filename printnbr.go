package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == -2147483648 {
		z01.PrintRune('-')
		z01.PrintRune('2')
		PrintNbr(147483648)
	} else if n < 0 {
		z01.PrintRune('-')
		PrintNbr(-n)
	} else if n >= 10 {
		PrintNbr(n / 10)
		z01.PrintRune(rune(n%10 + '0'))
	} else {
		z01.PrintRune(rune(n + '0'))
	}
}
