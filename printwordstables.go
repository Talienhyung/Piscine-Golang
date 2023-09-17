package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for i := range a {
		if len(a)-1 != i {
			printstr(a[i])
			z01.PrintRune('\n')
		} else {
			printstr(a[i])
		}
	}
}

func printstr(str string) {
	for _, i := range str {
		z01.PrintRune(i)
	}
}
