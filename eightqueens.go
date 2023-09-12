package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	for i := 15863724; i < 87654321; i++ {
		tab := transformation(i)
		if valide(tab) && verification(tab) {
			printTableau(tab)
		}
	}
}

func transformation(n int) []int {
	digits := make([]int, 0)
	t2 := make([]int, 0)
	for n > 0 {
		p := n % 10
		digits = append(digits, p)
		n = n / 10
	}
	for i := len(digits) - 1; i >= 0; i-- {
		t2 = append(t2, digits[i])
	}
	return t2
}

func valide(tab []int) bool {
	for i := range tab {
		if tab[i] == 0 || tab[i] == 9 {
			return false
		}
	}
	for i, c := range tab {
		for j, d := range tab {
			if i != j && c == d {
				return false
			}
		}
	}
	return true
}

func validationManger(i int, j int, s []int) string {
	if i+j < 8 {
		if s[i+j] == s[i]+j || s[i+j] == s[i]-j {
			return "false"
		}
	} else if i-j >= 0 {
		if s[i-j] == s[i]+j || s[i-j] == s[i]-j {
			return "false"
		}
	}
	return "true"
}

func verification(tab []int) bool {
	for i := 1; i < 8; i++ {
		for j := 1; j < 8; j++ {
			if validationManger(i, j, tab) == "false" {
				return false
			}
		}
	}
	return true
}

func printTableau(tab []int) {
	for i := 0; i < len(tab); i++ {
		z01.PrintRune(rune(tab[i] + '0'))
	}
	z01.PrintRune('\n')
}
