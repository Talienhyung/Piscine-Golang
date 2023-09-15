package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	var tab []rune
	args := os.Args[1:]
	for _, j := range args {
		for _, l := range j {
			if estVoyelle(l) {
				tab = append(tab, l)
			}
		}
	}
	newVoyelle := tabEnvers(tab)
	newPhrase := grandremplacement(rassemblement(args), newVoyelle)
	for i := range newPhrase {
		z01.PrintRune(rune(newPhrase[i]))
	}
	z01.PrintRune('\n')
}

func estVoyelle(s rune) bool {
	if s == 'a' || s == 'A' || s == 'e' || s == 'E' || s == 'i' || s == 'I' || s == 'o' || s == 'O' || s == 'u' || s == 'U' {
		return true
	} else {
		return false
	}
}

func tabEnvers(tab []rune) []string {
	var newTab []string
	for w := len(tab) - 1; w >= 0; w-- {
		newTab = append(newTab, string(tab[w]))
	}
	return newTab
}

func rassemblement(tab []string) string {
	phrase := ""
	for i, b := range tab {
		if i != len(tab) {
			phrase += b
			phrase += " "
		} else {
			phrase += b
		}

	}
	return phrase
}

func grandremplacement(phrase string, tab []string) string {
	i := 0
	newPhrase := ""
	for j := range phrase {
		if estVoyelle(rune(phrase[j])) {
			newPhrase += tab[i]
			i++
		} else {
			newPhrase += string(phrase[j])
		}
	}
	return newPhrase
}
