package main

import (
	"fmt"
	"os"
)

func main() {
	truc := "\n"
	arg := os.Args[1:]
	if len(arg) == 0 {
		return
	} else if arg[0] != "-c" {
		return
	}
	x := Atoi(arg[1])
	r := ""
	ok := false
	for i := 2; i < len(arg); i++ {
		if i == len(arg)-1 {
			truc = ""
		}
		fileName := arg[i]
		data, err := os.ReadFile(fileName)
		if err != nil {
			ok = true
			fmt.Printf("open " + fileName + ": no such file or directory" + truc + truc)
		} else {
			if len(data)-x < 0 {
				x = 0
			} else {
				x = len(data) - x
			}
			for i := x; i < len(data); i++ {
				r += string(data[i])
			}
			fmt.Printf("==> " + fileName + " <==\n" + r + truc)
		}
	}
	if ok {
		os.Exit(1)
	}
}

func Atoi(s string) int {
	non := 0
	nega := 2
	for i := range s {
		if s[i] != '-' && s[i] != '+' {
			if s[i] < '0' || s[i] > '9' {
				return 0
			}
		}
		if s[i] >= '0' && s[i] <= '9' {
			non = non*10 + int(s[i]-'0')
		}
		if s[i] == '-' {
			if i == 0 {
				nega = -1
			} else {
				return 0
			}
		} else if s[i] == '+' {
			if i == 0 {
				nega = 1
			} else {
				return 0
			}
		}
	}
	if nega == 2 {
		nega = 1
	}
	return non * nega
}
