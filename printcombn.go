package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n > 0 && n < 10 {
		var tableau []int
		for i := 0; i < n; i++ {
			tableau = append(tableau, i)
		}
		printDigits(tableau)
		for hasNext(tableau) {
			next(tableau)
			printDigits(tableau)
		}
	}
}

func printDigits(tableau []int) {
	for _, d := range tableau {
		z01.PrintRune(rune(d + '0'))
	}
	if tableau[0] != 10-len(tableau) {
		z01.PrintRune(',')
		z01.PrintRune(' ')
	} else {
		z01.PrintRune('\n')
	}
}

func hasNext(tableau []int) bool {
	n := len(tableau)
	for i := 0; i < n; i++ {
		if tableau[i] != 10-n+i {
			return true
		}
	}
	return false
}

func next(tableau []int) {
	n := len(tableau)
	i := n - 1
	for i >= 0 && tableau[i] == 10-n+i {
		i--
	}
	if i >= 0 {
		tableau[i]++
		for j := i + 1; j < n; j++ {
			tableau[j] = tableau[j-1] + 1
		}
	}
}
