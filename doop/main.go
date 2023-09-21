package main

import (
	"os"
)

func main() {
	args := os.Args[1:] // Récupérer tous les arguments en ligne de commande
	if len(args) != 3 {
		return
	}
	ch1 := stringToInt(args[0])
	ch2 := stringToInt(args[2])
	if ch1 == -1 || ch2 == -1 || ch2 > 9223372036854775807 || ch1 > 9223372036854775807 || ch1 < -9223372036854775808 || ch2 < -9223372036854775808 {
		return
	}
	switch args[1] {
	case "+":
		resultat := multpicaption(ch1, args[1], ch2)
		resutat2 := intToString(resultat)
		if resultat >= 9223372036854775805 || resultat < -9223372036854775808 {
			return
		} else if resutat2 == "-" {
			return
		}
		os.Stdout.WriteString(resutat2 + "\n")
	case "-":
		resultat := multpicaption(ch1, args[1], ch2)
		resutat2 := intToString(resultat)
		if resultat >= 9223372036854775805 || resultat < -9223372036854775808 {
			return
		} else if resutat2 == "-" {
			return
		}
		os.Stdout.WriteString(resutat2 + "\n")
	case "*":
		resultat := multpicaption(ch1, args[1], ch2)
		resutat2 := intToString(resultat)
		if resultat >= 9223372036854775805 || resultat < -9223372036854775808 {
			return
		} else if resutat2 == "-" {
			return
		}
		os.Stdout.WriteString(resutat2 + "\n")
	case "/":
		if ch2 == 0 {
			os.Stdout.WriteString("No division by 0\n")
		} else {
			resultat := multpicaption(ch1, args[1], ch2)
			resutat2 := intToString(resultat)
			if resultat >= 9223372036854775805 || resultat < -9223372036854775808 {
				return
			} else if resutat2 == "-" {
				return
			}
			os.Stdout.WriteString(resutat2 + "\n")
		}
	case "%":
		if ch2 == 0 {
			os.Stdout.WriteString("No modulo by 0\n")
		} else {
			resultat := multpicaption(ch1, args[1], ch2)
			resutat2 := intToString(resultat)
			if resultat >= 9223372036854775805 || resultat < -9223372036854775808 {
				return
			} else if resutat2 == "-" {
				return
			}
			os.Stdout.WriteString(resutat2 + "\n")
		}
	default:
		return
	}
}

func stringToInt(s string) int {
	into := 0
	nega := 1
	for j, i := range s {
		if i < '0' || i > '9' {
			if i == '-' && j == 0 || i == '+' && j == 0 {
				nega = -1
			} else {
				return -1
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
