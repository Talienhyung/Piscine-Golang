// package main

// import (
// 	"os"
// )

// func main() {
// 	args := os.Args[:1]
// 	if len(args) > 3 || len(args) < 3 {
// 		return
// 	} else if args[1] != "+" && args[1] != "-" && args[1] != "/" && args[1] != "*" && args[1] != "%" {
// 		return
// 	} else if StringtoInt(args[0]) >= 9223372036854775807 || StringtoInt(args[2]) >= 9223372036854775807 || StringtoInt(args[0]) <= -9223372036854775808 || StringtoInt(args[2]) <= -9223372036854775808 {
// 		os.Stdout.WriteString("No division by 0")
// 	} else if args[1] == "/" && args[2] == "0" {
// 		os.Stdout.WriteString("No division by 0")
// 	} else if args[1] == "%" && args[2] == "0" {
// 		os.Stdout.WriteString("No modulo by 0")
// 	} else {
// 		resultat := multpicaption(StringtoInt(args[0]), args[1], StringtoInt(args[2]))
// 		resutat2 := intToString(resultat)
// 		os.Stdout.WriteString(resutat2)
// 	}
// }

// func StringtoInt(s string) int {
// 	into := 0
// 	for _, i := range s {
// 		into = into*10 + int(i-'0')
// 	}
// 	return into
// }

// func intToString(n int) string {
// 	if n == 0 {
// 		return "0"
// 	}
// 	sign := ""
// 	if n < 0 {
// 		sign = "-"
// 		n = -n
// 	}
// 	result := ""
// 	for n > 0 {
// 		digit := n % 10
// 		result = string('0'+digit) + result
// 		n /= 10
// 	}
// 	return sign + result
// }

// func multpicaption(ch1 int, div string, ch2 int) int {
// 	if div == "+" {
// 		return ch1 + ch2
// 	} else if div == "-" {
// 		return ch1 - ch2
// 	} else if div == "/" {
// 		return ch1 / ch2
// 	} else if div == "*" {
// 		return ch1 * ch2
// 	} else if div == "%" {
// 		return ch1 % ch2
// 	} else {
// 		return 0
// 	}
// }

package main

import (
	"os"
)

func main() {
	args := os.Args[1:] // Récupérer tous les arguments en ligne de commande
	if len(args) != 3 {
		os.Stdout.WriteString("Usage: calculatrice <nombre1> <opérateur> <nombre2>\n")
		return
	}
	ch1 := stringToInt(args[0])
	ch2 := stringToInt(args[2])
	if ch1 == -1 || ch2 == -1 {
		os.Stdout.WriteString("Erreur lors de la conversion des nombres\n")
		return
	}
	switch args[1] {
	case "+":
		resultat := multpicaption(ch1, args[1], ch2)
		resutat2 := intToString(resultat)
		os.Stdout.WriteString(resutat2 + "\n")
	case "-":
		resultat := multpicaption(ch1, args[1], ch2)
		resutat2 := intToString(resultat)
		os.Stdout.WriteString(resutat2 + "\n")
	case "*":
		resultat := multpicaption(ch1, args[1], ch2)
		resutat2 := intToString(resultat)
		os.Stdout.WriteString(resutat2 + "\n")
	case "/":
		if ch2 == 0 {
			os.Stdout.WriteString("No division by 0\n")
		} else {
			resultat := multpicaption(ch1, args[1], ch2)
			resutat2 := intToString(resultat)
			os.Stdout.WriteString(resutat2)
		}
	case "%":
		if ch2 == 0 {
			os.Stdout.WriteString("No modulo by 0\n")
		} else {
			resultat := multpicaption(ch1, args[1], ch2)
			resutat2 := intToString(resultat)
			os.Stdout.WriteString(resutat2 + "\n")
		}
	default:
		os.Stdout.WriteString("Opérateur non valide\n")
	}
}

func stringToInt(s string) int {
	into := 0
	nega := 1

	for _, i := range s {
		if i < '0' || i > '9' {
			if string(s[0]) == "-" {
				nega = -1
			}
		} else {
			into = into*10 + int(i-'0')
		}
	}
	return into * nega
}

func intToString(n int) string {
	if n == 0 {
		return "0"
	}
	sign := ""
	if n < 0 {
		sign = "-"
		n = -n
	}
	result := ""
	for n > 0 {
		digit := n % 10
		result = string('0'+digit) + result
		n /= 10
	}
	return sign + result
}

func multpicaption(ch1 int, div string, ch2 int) int {
	if div == "+" {
		return ch1 + ch2
	} else if div == "-" {
		return ch1 - ch2
	} else if div == "/" {
		return ch1 / ch2
	} else if div == "*" {
		return ch1 * ch2
	} else if div == "%" {
		return ch1 % ch2
	} else {
		return 0
	}
}
