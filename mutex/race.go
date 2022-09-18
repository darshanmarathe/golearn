package main

import (
	"fmt"
	"time"
)

// This is an example of race condition
// 2 goroutines tries to read&write sharedInt and there is no access control.
// go run -race rage.go
// https://medium.com/trendyol-tech/race-conditions-in-golang-511314c0b85
var sharedInt int = 0
var unusedValue int = 0

func runSimpleReader() {

	for {
		var val int = sharedInt
		if val%10 == 0 {
			unusedValue = unusedValue + 1
		}
	}
}

func runSimpleWriter() {
	for {
		sharedInt = sharedInt + 1
	}
}

func main() {
	startSimpleReadWrite()
}

func startSimpleReadWrite() {
	go runSimpleReader()
	go runSimpleWriter()
	fmt.Print("stated....")
	time.Sleep(10 * time.Second)
}
