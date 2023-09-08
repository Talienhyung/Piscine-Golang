package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	digits := make([]int, 0)
	for n > 0 {
		p := n % 10
		digits = append(digits, p)
		n = n / 10
	}
	for i := 0; i < len(digits)-1; i++ {
		for j := 0; j < len(digits)-i-1; j++ {
			if digits[j] > digits[j+1] {
				digits[j], digits[j+1] = digits[j+1], digits[j]
			}
		}
	}
	for i := 0; i < len(digits); i++ {
		z01.PrintRune(rune(digits[i] + '0'))
	}
}
