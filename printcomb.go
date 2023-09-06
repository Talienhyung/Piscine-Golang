package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := 0; i <= 10; i++ {
		for j := 0; j <= 10; j++ {
			for k := 0; k <= 10; k++ {
				if i < j && j < k {
					z01.PrintRune(rune(i))
					z01.PrintRune(rune(i))
					z01.PrintRune(rune(i))
					if i != 7 || j != 8 || k != 9 {
						z01.PrintRune(rune(','))
						z01.PrintRune(rune(' '))
					}
				}
			}
		}
	}
}
