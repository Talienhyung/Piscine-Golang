package main

import "fmt"

var Grille = [10][10]rune{
	{'a', 'v', 't', 'g', 'e', 'r', 'e', 'c', 'c', 'a'},
	{'l', 'a', 'o', 'a', 'b', 'o', 'm', 'h', 'c', 'e'},
	{'e', 't', 'u', 'l', 'e', 'm', 'o', 'a', 'r', 'e'},
	{'m', 'a', 't', 'e', 'e', 'a', 's', 'i', 'e', 'f'},
	{'a', 'r', 'a', 'd', 'i', 'u', 't', 's', 'u', 'r'},
	{'c', 'h', 't', 'i', 'r', 'a', 'r', 'e', 'e', 'a'},
	{'a', 'a', 'r', 'o', 'm', 'r', 'n', 'n', 'l', 'n'},
	{'j', 'e', 'r', 'm', 'i', 'e', 't', 'o', 'u', 't'},
	{'o', 'b', 's', 'e', 'c', 'o', 'o', 'l', 'e', 'l'},
	{'u', 't', 'r', 'e', 's', 'c', 'h', 'i', 'e', 'n'},
}

var MotATrouver = "el"

func main() {
	runes := []rune(MotATrouver)
	found := false
	for i := 0; i < len(Grille); i++ { // axe colone
		for t := 0; t < len(Grille[i]); t++ { // axe ligne
			if t+len(runes) <= len(Grille[i]) {
				if Grille[i][t] == runes[0] {
					for n := 1; n < len(runes); n++ {
						if i-n < 0 || t+n > 10 {
							found = false
							break
						}
						if Grille[i-n][t+n] == runes[n] {
							found = true
						} else {
							found = false
							break
						}
					}
					if found {
						fmt.Println("tre")
						return
					}
				}
			}
		}
	}
	fmt.Println("False")
}
