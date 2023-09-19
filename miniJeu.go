package piscine

import (
	"fmt"
	"math/rand"
)

var inputs string

func MenuMiniJeu() {
	fmt.Println("\n-----------------------------------------")
	fmt.Println("Choisir un jeu parmi les jeux suivants :")
	fmt.Println("1 > Pierre Feuille Ciseaux")
	fmt.Println("2 > Trouver le nombre")
	fmt.Println("3 > Morpion")
	fmt.Println("4 > Quitter")
	fmt.Println("-----------------------------------------")
	inputs = input("\nChoix du jeu : ", inputs)
	if inputs == "1" {
		Pierre()
	} else if inputs == "2" {
		Nombre()
	} else if inputs == "3" {
		Morpion()
	} else if inputs == "4" {
		return
	}
}

func Pierre() {
	scoreOrdi := 0
	scoreUti := 0
	fmt.Println("\n-----------------------------------------\nJouons à pierre feuille ciseaux !!\n-----------------------------------------")
	for scoreOrdi < 5 && scoreUti < 5 {
		choixOrdi := random(1, 3)
		choixUti := input("\nPierre, Feuille ou Ciseau", inputs)
		if choixOrdi == 1 && choixUti == "pierre" {
			fmt.Println("Vous jouez pierre et l'ordinateur aussi !")
			fmt.Printf("Ordi : %d Vous : %d \n", scoreOrdi, scoreUti)
		}
		if choixOrdi == 2 && choixUti == "pierre" {
			fmt.Println("Vous jouez pierre et l'ordinateur feuille !")
			scoreOrdi++
			fmt.Printf("Ordi : %d Vous : %d \n", scoreOrdi, scoreUti)
		}
		if choixOrdi == 3 && choixUti == "pierre" {
			fmt.Println("Vous jouez pierre et l'ordinateur ciseaux !")
			scoreUti++
			fmt.Printf("Ordi : %d Vous : %d \n", scoreOrdi, scoreUti)
		}
		if choixOrdi == 1 && choixUti == "feuille" {
			fmt.Println("Vous jouez feuille et l'ordinateur pierre !")
			scoreUti++
			fmt.Printf("Ordi : %d Vous : %d \n", scoreOrdi, scoreUti)
		}
		if choixOrdi == 2 && choixUti == "feuille" {
			fmt.Println("Vous jouez feuille et l'ordinateur aussi !")
			fmt.Printf("Ordi : %d Vous : %d \n", scoreOrdi, scoreUti)
		}
		if choixOrdi == 3 && choixUti == "feuille" {
			fmt.Println("Vous jouez feuille et l'ordinateur ciseaux !")
			scoreOrdi++
			fmt.Printf("Ordi : %d Vous : %d \n", scoreOrdi, scoreUti)
		}
		if choixOrdi == 1 && choixUti == "ciseaux" {
			fmt.Println("Vous jouez ciseaux et l'ordinateur pierre !")
			scoreOrdi++
			fmt.Printf("Ordi : %d Vous : %d \n", scoreOrdi, scoreUti)
		}
		if choixOrdi == 2 && choixUti == "ciseaux" {
			fmt.Println("Vous jouez ciseaux et l'ordinateur feuille !")
			scoreUti++
			fmt.Printf("Ordi : %d Vous : %d \n", scoreOrdi, scoreUti)
		}
		if choixOrdi == 3 && choixUti == "ciseaux" {
			fmt.Println("Vous jouez ciseaux et l'ordinateur aussi !")
			fmt.Printf("Ordi : %d Vous : %d \n", scoreOrdi, scoreUti)
		}
	}
	if scoreOrdi == 5 {
		fmt.Println("\n L'ordinateur gagne ! (Sombre merde)")
	}
	if scoreUti == 5 {
		fmt.Println("\nouiTu gagne !! Bravo immense BG !!!")
	}
	if input("\n-----------------------------------------\nVoulez-vous rejouer ?\n-----------------------------------------\n", inputs) == "oui" {
		Pierre()
	} else {
		MenuMiniJeu()
	}
}

func input(s string, input string) string {
	fmt.Println(s)
	fmt.Scanln(&input)
	return input
}

func random(x, y int) int {
	randomNumber := rand.Intn(y-x+1) + x
	return randomNumber
}

func stringtoint(s string) int {
	runes := []rune(s)
	n := 0
	for i := range runes {
		n = n*10 + int(s[i]-'0')
	}
	return n
}

func Nombre() {
	gagner := false
	fmt.Printf("\n-----------------------------------------\nTrouver un chiffre entre %d et %d !!\n-----------------------------------------\n", 0, 500)
	chiffreAtrouver := random(0, 500)
	for !gagner {
		chiffre := stringtoint(input("\nPropose un chiffre :", inputs))
		if chiffre == chiffreAtrouver {
			fmt.Println("\nENFIN BRAVO !!")
			gagner = true
		}
		if chiffre > chiffreAtrouver {
			fmt.Println("\nTrop grand !!")
		}
		if chiffre < chiffreAtrouver {
			fmt.Println("\nTrop petit !!")
		}
	}
	if input("\n-----------------------------------------\nVoulez-vous rejouer ?\n-----------------------------------------\n", inputs) == "oui" {
		Nombre()
	} else {
		MenuMiniJeu()
	}
}

func Morpion() {
	tab := make([]string, 9)
	gagner := false
	utilisateur := "X"
	tab = premiereTab(tab)
	printTab(tab)
	for !gagner {
		choixUti := choixUtilisateur(utilisateur)
		if !verificationChoix(choixUti, tab) {
			fmt.Println("Mauvais choix. Si tu perds c'est ton probleme pas le miens.")
		} else {
			tab = nouveauTableau(choixUti, tab, utilisateur)
			printTab(tab)
		}
		if quiGagne(utilisateur, tab) {
			gagner = true
		} else {
			utilisateur = changementUti(utilisateur)
		}

	}
	fmt.Println("\nBRAVO " + utilisateur + " TU ES UN IMMENSE BG !!")
	if input("\n-----------------------------------------\nVoulez-vous rejouer ?\n-----------------------------------------\n", inputs) == "oui" {
		Morpion()
	} else {
		MenuMiniJeu()
	}
}

func verificationChoix(choix string, tab []string) bool {
	if tab[stringtoint(choix)-1] == "0" || tab[stringtoint(choix)-1] == "X" {
		return false
	} else {
		return true
	}
}

func nouveauTableau(choixUti string, tab []string, utilisateur string) []string {
	tab[stringtoint(choixUti)-1] = utilisateur
	return tab
}

func changementUti(utili string) string {
	if utili == "X" {
		utili = "0"
	} else {
		utili = "X"
	}
	return utili
}

func choixUtilisateur(uti string) string {
	inputs = input("\n-----------------------------------------\n"+uti+" : Où voulez-vous jouer ?", inputs)
	return inputs
}

func quiGagne(uti string, tab []string) bool {
	if tab[0] == uti && tab[1] == uti && tab[2] == uti {
		return true
	} else if tab[3] == uti && tab[4] == uti && tab[5] == uti {
		return true
	} else if tab[6] == uti && tab[7] == uti && tab[8] == uti {
		return true
	} else if tab[0] == uti && tab[4] == uti && tab[8] == uti {
		return true
	} else if tab[6] == uti && tab[4] == uti && tab[2] == uti {
		return true
	} else if tab[6] == uti && tab[3] == uti && tab[0] == uti {
		return true
	} else if tab[7] == uti && tab[4] == uti && tab[1] == uti {
		return true
	} else if tab[8] == uti && tab[5] == uti && tab[2] == uti {
		return true
	}
	return false
}

func printTab(tab []string) {
	fmt.Println("\n" + tab[6] + " | " + tab[7] + " | " + tab[8] + "\n---------\n" + tab[3] + " | " + tab[4] + " | " + tab[5] + "\n---------\n" + tab[0] + " | " + tab[1] + " | " + tab[2])
}

func premiereTab(tab []string) []string {
	for i := 0; i <= 8; i++ {
		tab[i] = string(rune(i + 1 + '0'))
	}
	return tab
}
