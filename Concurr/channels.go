package main

import (
	"fmt"
	"runtime"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU() * 2)

	ch := make(chan string)
	doneCh := make(chan bool)
	go abcgen(ch)
	go abcgen(ch)
	go abcgen(ch)
	go Printer(ch, doneCh)
	fmt.Println("The number of GO 	Cores:", runtime.NumGoroutine())

	println("here it came first")
	<-doneCh
}

func abcgen(ch chan string) {
	for i := byte('a'); i <= byte('z'); i++ {
		ch <- string(i)
	}
	close(ch)
}

func Printer(ch chan string, doneCh chan bool) {
	for i := range ch {
		println(i)
	}
	doneCh <- true
}
