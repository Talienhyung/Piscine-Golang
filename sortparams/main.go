package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Récupérer les arguments de la ligne de commande
	args := os.Args[1:]

	// Trier les arguments en ordre ASCII
	sortArgs(args)

	// Imprimer les arguments triés
	for _, arg := range args {
		for _, char := range arg {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}

func sortArgs(args []string) {
	for i := 0; i < len(args); i++ {
		for j := i + 1; j < len(args); j++ {
			if compareStrings(args[i], args[j]) > 0 {
				args[i], args[j] = args[j], args[i]
			}
		}
	}
}

func compareStrings(a, b string) int {
	i := 0
	for i < len(a) && i < len(b) {
		if a[i] != b[i] {
			return int(a[i]) - int(b[i])
		}
		i++
	}
	return len(a) - len(b)
}
