package main

import "fmt"

type WARN string
type INFO string
type ERROR string
type FIX string
type FATAL string

type LOG interface {
	String() string
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

func (s FATAL) String() string {
	return string("FATAL:: " + s)
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

	warn := WARN("Hello")
	_error := ERROR("Oh no not again !!!")
	fix := FIX("To Fix this download the latest version")
	info := INFO(fmt.Sprintf("str type: %T", warn))

	fatal := FATAL("Ohhh this is bad i am shutting down")

	printInfo(info)

	printInfo(warn)
	printInfo(_error)
	printInfo(fix)

	printInfo(fatal)
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
		fmt.Println(string(colorCyan), v.String())
	}
}
