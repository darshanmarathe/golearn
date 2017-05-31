package main

import "time"
import "runtime"

func main() {
	runtime.GOMAXPROCS(10)
	go abcgen()
	time.Sleep(100 * time.Millisecond)

}

func abcgen() {
	for i := byte('a'); i <= byte('z'); i++ {
		go println(string(i))
	}
}
