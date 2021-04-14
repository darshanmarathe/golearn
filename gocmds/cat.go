package main

import (
	"fmt"
	"os"
)

func main() {

	Args := os.Args[1:]
	if len(Args) == 0 {
		printHelpMessage()
		return
	}
	for _, v := range Args {
		fmt.Println(v)
	}
}

func printHelpMessage() {
	fmt.Println("Looks like you did not use program correctly.")
	fmt.Println("Try using program like.")
	fmt.Println("cat filename.txt will print the file")
	fmt.Println("cat filename.txt -l will print the file with lines")
	fmt.Println("cat filename.txt -c will print and copy the content of the file to memory")
}
