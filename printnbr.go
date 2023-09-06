package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	var digits []rune
	digits = make([]rune, 3)
	for n > 0 {
		digits = append(digits, rune(n%10+'0'))
		n = n / 10
	}
	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(digits[i])
	}
}
