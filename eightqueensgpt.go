package piscine

import "github.com/01-edu/z01"

// EightQueens est la fonction principale qui trouve et imprime toutes les solutions possibles au problème des huit reines.
func EightQueensgpt() {
	var tab [8]int // Un tableau statique pour stocker les positions des reines dans chaque rangée.
	placeQueensgpt(0, tab)
}

// placeQueensgpt est une fonction récursive qui place les reines dans chaque rangée.
func placeQueensgpt(row int, tab [8]int) {
	if row == 8 {
		// Si nous avons placé toutes les huit reines, nous imprimons la solution.
		printTableaugpt(tab)
		return
	}

	for col := 0; col < 8; col++ {
		tab[row] = col // Essayez de placer une reine dans cette colonne.

		// Vérifiez si la configuration actuelle est validegpt.
		if validegpt(tab, row) {
			// Si la configuration est validegpt, passez à la rangée suivante.
			placeQueensgpt(row+1, tab)
		}
	}
}

// validegpt vérifie si une configuration donnée des reines est validegpt.
func validegpt(tab [8]int, row int) bool {
	for i := 0; i < row; i++ {
		// Vérifiez si deux reines partagent la même colonne ou diagonale.
		if tab[i] == tab[row] || row-i == absgpt(tab[row]-tab[i]) {
			return false
		}
	}
	return true
}

// printTableaugpt imprime une configuration de reines.
func printTableaugpt(tab [8]int) {
	for _, val := range tab {
		z01.PrintRune(rune(val + '0'))
	}
	z01.PrintRune('\n')
}

// absgpt retourne la valeur absgptolue d'un entier.
func absgpt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
