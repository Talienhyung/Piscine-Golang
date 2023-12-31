package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) string {
	no := ""
	if len(base) < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return "o"
	}
	for i, c := range base {
		for j, d := range base {
			if i != j && c == d {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return "o"
			}
		}
	}
	for _, c := range base {
		if c == '+' || c == '-' {
			z01.PrintRune('N')
			z01.PrintRune('V')
			return "o"
		}
	}

	length := len(base)
	sign := 1
	if nbr < 0 {
		sign = -1
	}
	digits := []rune{}

	for nbr < 0 {
		digits = append(digits, rune(base[nbr%length*-1]))
		nbr /= length
	}

	for nbr > 0 {
		digits = append(digits, rune(base[nbr%length]))
		nbr /= length
	}

	if sign == -1 {
		z01.PrintRune('-')
	}

	for i := len(digits) - 1; i >= 0; i-- {
		no += string(digits[i])
	}
	return no
}
