package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for i := 0; i <= 10; i++ {
			time.Sleep(500 * time.Millisecond)
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for i := 0; i <= 15; i++ {

			time.Sleep(200 * time.Millisecond)
			ch2 <- i
		}
		close(ch2)
	}()
	chan1Open, chan2Open := false, false
	v := 0
	for {
		chan1Open, chan2Open = true, true
		select {
		case v, chan1Open = <-ch1:
			if chan1Open {
				fmt.Println(v, "chan1")
			}
		default:
		}
		select {
		case v, chan2Open = <-ch2:
			if chan2Open {
				fmt.Println(v, "chan2")
			}
		default:
		}
		if chan1Open == false && chan2Open == false {
			fmt.Println("closing.....")
			return

		}

	}

}
