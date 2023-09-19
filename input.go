package piscine

import "fmt"

func Input() {
	var input string
	fmt.Print("Nouveau J: ")
	fmt.Scanln(&input)
	fmt.Println(input)
}
