package main

import "fmt"

func pingfunc(pings chan<- string, msg string) {
	pings <- msg
}

func pongfunc(pings <-chan string, pongs chan<- string) {
	// msg := <-pings
	// pongs <- msg

	pongs <- <-pings
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	pingfunc(pings, "passed message")
	pongfunc(pings, pongs)

	fmt.Println(<-pongs)
	fmt.Println(<-pings)
}
