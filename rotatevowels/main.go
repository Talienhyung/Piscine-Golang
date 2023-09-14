package main

import (
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		println()
		return
	}

	for i, arg := range args {
		mirroredArg := mirrorVowels(arg)
		print(mirroredArg)
		if i < len(args)-1 {
			print(" ")
		}
	}
	println()
}

func mirrorVowels(s string) string {
	sRunes := []rune(s)

	for i := 0; i < len(sRunes)/2; i++ {
		if isVowel(sRunes[i]) && isVowel(sRunes[len(sRunes)-1-i]) {
			sRunes[i], sRunes[len(sRunes)-1-i] = sRunes[len(sRunes)-1-i], sRunes[i]
		}
	}

	return string(sRunes)
}

func isVowel(r rune) bool {
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' || r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U'
}
