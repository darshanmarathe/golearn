package main

import (
	"fmt"
)

type PrinterFunc func(string) ()

func main() {
	Concatnitor := Concat("Hello")
	Concatnitor(" World")
	Concatnitor(" this")
	Concatnitor(" is")
	Concatnitor(" go")
	Concatnitor(" language")
	Concatnitor(" showing closures")
}

func Concat(s string) (func(ar string)) {

	var line = s
	return func(z string) {
		line += z
		fmt.Println(line)
	}

}
