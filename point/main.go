package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	setPoint(points)
	printRune('x')
	printRune('=')
	printInt(points.x)
	printRune(',')
	printRune(' ')
	printRune('y')
	printRune('=')
	printInt(points.y)
	printRune('\n')
}

func printRune(r rune) {
	z01.PrintRune(r)
}

func printInt(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	if n >= 10 {
		printInt(n / 10)
	}
	z01.PrintRune(rune(n%10) + '0')
}
