package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	if n < 0 {
		z01.PrintRune('-')
	}

	digits := make([]int, 0)
	for n > 0 {
		p := n % 10
		digits = append(digits, p)
		n = n / 10
	}
	for n < 0 {
		p := n % 10
		digits = append(digits, -p)
		n = n / 10
	}

	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(rune(digits[i] + '0'))
	}
}
