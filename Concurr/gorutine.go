package main

import (
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(10)
	abcgen()
	time.Sleep(100 * time.Millisecond)

}

func abcgen() {
	for i := byte('a'); i <= byte('z'); i++ {
		go println(string(i))
	}
}
