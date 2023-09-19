package main

import (
	"os"
)

func main() {
	args := os.Args[:1]
	if len(args) > 3 || len(args) < 3 {
		return
	} else if args[1] != "+" && args[1] != "-" && args[1] != "/" && args[1] != "*" && args[1] != "%" {
		return
	} else if args[0] >= "9223372036854775807" || args[2] >= "9223372036854775807" || args[0] <= "-9223372036854775809" || args[2] <= "-9223372036854775809" {
		return
	} else if args[1] == "/" && args[2] == "0" {
		os.Stdout.WriteString("No division by 0")
	} else if args[1] == "%" && args[2] == "0" {
		os.Stdout.WriteString("No modulo by 0")
	} else {
		resultat := multpicaption(StringtoInt(args[0]), args[1], StringtoInt(args[2]))
		resutat2 := intToString(resultat)
		os.Stdout.WriteString(resutat2)
	}
}

func StringtoInt(s string) int {
	into := 0
	for _, i := range s {
		into = into*10 + int(i-'0')
	}
	return into
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
