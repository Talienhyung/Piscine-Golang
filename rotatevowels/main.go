package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		z01.PrintRune('\n')
		return
	}
	for i, arg := range args {
		mirroredArg := mirrorVowels(arg)
		for _, r := range mirroredArg {
			z01.PrintRune(r)
		}
		if i < len(args)-1 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}

func mirrorVowels(s string) string {
	sRunes := []rune(s)
	for i, j := 0, len(sRunes)-1; i < j; {
		if isVowel(sRunes[i]) && isVowel(sRunes[j]) {
			sRunes[i], sRunes[j] = sRunes[j], sRunes[i]
			i++
			j--
		} else if !isVowel(sRunes[i]) {
			i++
		} else if !isVowel(sRunes[j]) {
			j--
		}
	}
	return string(sRunes)
}

func isVowel(r rune) bool {
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' || r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U'
}
