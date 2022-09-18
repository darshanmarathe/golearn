package main

import "fmt"

type WARN string
type INFO string
type ERROR string
type FIX string

type LOG interface {
	String() string
}

var colorReset = "\033[0m"
var colorRed = "\033[31m"
var colorGreen = "\033[32m"
var colorYellow = "\033[33m"
var colorBlue = "\033[34m"
var colorPurple = "\033[35m"
var colorCyan = "\033[36m"
var colorWhite = "\033[37m"

func main() {
	var warn LOG
	warn = WARN("Hello")

	error := ERROR("Oh no not again !!!")
	fmt.Printf("str type: %T\n", warn)
	printInfo(warn)
	printInfo(error)
}

func printInfo(msg LOG) {
	switch v := msg.(type) {
	case INFO:
		fmt.Println(string(colorWhite), v.String())
	case WARN:
		fmt.Println(string(colorYellow), v.String())
	case ERROR:
		fmt.Println(string(colorRed), v.String())
	case FIX:
		fmt.Println(string(colorGreen), v.String())
	default:
		fmt.Println(v.String())
	}
}

func (s WARN) String() string {
	return string("WARNING:: " + s)
}

func (s INFO) String() string {
	return string("INFO:: " + s)
}

func (s ERROR) String() string {
	return string("ERROR:: " + s)
}

func (s FIX) String() string {
	return string("FIX:: " + s)
}
