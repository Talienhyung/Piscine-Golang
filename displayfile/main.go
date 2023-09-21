package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1:]
	if len(arg) > 1 {
		fmt.Print("Too many arguments\n")
		return
	} else if len(arg) == 0 {
		fmt.Print("File name missing\n")
		return
	}
	fileName := arg[0]
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Print(string(data))
}
