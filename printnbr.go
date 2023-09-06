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
		if digits[i] == 0 {
			z01.PrintRune('0')
		}
		if digits[i] == 1 {
			z01.PrintRune('1')
		}
		if digits[i] == 2 {
			z01.PrintRune('2')
		}
		if digits[i] == 3 {
			z01.PrintRune('3')
		}
		if digits[i] == 4 {
			z01.PrintRune('4')
		}
		if digits[i] == 5 {
			z01.PrintRune('5')
		}
		if digits[i] == 6 {
			z01.PrintRune('6')
		}
		if digits[i] == 7 {
			z01.PrintRune('7')
		}
		if digits[i] == 8 {
			z01.PrintRune('8')
		}
		if digits[i] == 9 {
			z01.PrintRune('9')
		}
	}
}
