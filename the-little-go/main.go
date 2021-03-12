package main

import (
	"fmt"
	"os"
)

func myFunc() (string, int) {
	return "and a cool number is", 1337
}

func main() {
	test := ""

	if len(os.Args) != 2 {
		fmt.Println("Expected: go run main.go <Name>")
		os.Exit(1)
	}
	fmt.Println("My name is", os.Args[1], "and the file that I am executing is", os.Args[0], test)
	test, age := myFunc()
	fmt.Println("My name is", os.Args[1], "and the file that I am executing is", os.Args[0], test, age)

}
