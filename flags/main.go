package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	var list [2]string
	args := os.Args[1:]
	order := false
	if len(args) == 0 || args[0] == "--help" || args[0] == "-h" {
		fmt.Println("--insert")
		fmt.Println("  -i")
		fmt.Println("	This flag inserts the string into the string passed as argument.")
		fmt.Println("--order")
		fmt.Println("  -o")
		fmt.Println("	This flag will behave like a boolean, if it is called it will order the argument.")
	}
	list[1] = isinsert(args[0])
	args = args[1:]
	if args[0] == "--order" || args[0] == "-o" {
		args = args[1:]
		order = true
	}
	if len(args) != 0 {
		list[0] = args[0]
	}
	l := list[0] + list[1]
	newlist := []rune(l)
	if order {
		SortArgs(newlist)
	}
	for _, char := range newlist {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}

func SortArgs(args []rune) {
	for i := 0; i < len(args); i++ {
		for j := i + 1; j < len(args); j++ {
			if args[i] < args[j] {
				args[i], args[j] = args[j], args[i]
			}
		}
	}
}

func isinsert(arg string) string {
	new := ""
	for j := 0; j < len(arg); j++ {
		if string(arg[j]) == "=" {
			for i := 1; i <= len(arg); i++ {
				new += string(arg[i])
			}
			return new
		}
	}
	return arg
}
