package main

import (
	"fmt"
	"os"
)

func Doop() {
	args := os.Args[:1]
	if len(args) > 3 || len(args) < 3 {
		return
	} else if args[1] != "+" && args[1] != "-" && args[1] != "/" && args[1] != "*" && args[1] != "%" {
		return
	} else if args[0] >= "9223372036854775807" || args[2] >= "9223372036854775807" || args[0] <= "-9223372036854775809" || args[2] <= "-9223372036854775809" {
		return
	} else if args[1] == "/" && args[2] == "0" {
		fmt.Println("No division by 0")
	} else if args[1] == "%" && args[2] == "0" {
		fmt.Println("No modulo by 0")
	} else {
		os.Stdout.WriteString(string(rune(multpicaption(StringtoInt(args[0]), args[1], StringtoInt(args[2]))) + '0'))
	}
}

func StringtoInt(s string) int {
	into := 0
	for _, i := range s {
		into = into*10 + int(i-'0')
	}
	return into
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
