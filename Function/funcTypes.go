package main

import (
	"fmt"
)

type PrinterFunc func(string) ()

func main() {
	Concat(Print, "yes", "golang", "is", "a", "functional", "programing", "language")
	Reverse(Print, "yes", "golang", "is", "a", "functional", "programing", "language")
	Concat(func(line string) {
		fmt.Println(line)
	}, "Anonomus", "functions", "?")
}

func Concat(onDone func(string), messages ...string) {
	line := ""
	for _, later := range messages {
		line += " " + later
	}
	onDone(line)
}

func Reverse(onDone PrinterFunc, meesages ...string) {
	line := ""
	max := len(meesages)

	for index := 0; index < len(meesages); index++ {
		max--
		line += " " + meesages[max]

	}
	onDone(line)

}

func Print(msg string) {
	fmt.Println(msg)
}
