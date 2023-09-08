package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if len(base) < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	for i, c := range base {
		for j, d := range base {
			if i != j && c == d {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			}
		}
	}

	for _, c := range base {
		if c == '+' || c == '-' {
			z01.PrintRune('N')
			z01.PrintRune('V')
			return
		}
	}

	length := len(base)
	sign := 1
	if nbr < 0 {
		sign = -1
		nbr = -nbr
	}

	digits := []rune{}
	for nbr > 0 {
		digits = append(digits, rune(base[nbr%length]))
		nbr /= length
	}

	if sign == -1 {
		z01.PrintRune('-')
	}

	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(digits[i])
	}
}
