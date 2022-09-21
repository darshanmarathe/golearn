package main

import (
	"fmt"
	"time"
)

func main() {

	messages := make(chan string)
	messages2 := make(chan string)

	go func() { messages <- "ping" }()
	go func() {
		time.Sleep(5 * time.Second)
		messages2 <- "pong"
	}()

	go func() {
		time.Sleep(10 * time.Second)
		messages2 <- "pong ponga"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		messages2 <- "pongoooo"

		time.Sleep(1 * time.Second)
		messages2 <- "bango bango"

		time.Sleep(1 * time.Second)
		messages2 <- "bango bango bango"

		time.Sleep(1 * time.Second)
		messages2 <- "bango bango bango bango"
	}()

	go func() {
		for i := 0; i < 20; i++ {
			time.Sleep(1 * time.Second)
			messages2 <- "Bango" + fmt.Sprintf("%d", i)
		}
		close(messages2)
	}()

	msg := <-messages

	fmt.Println(msg)

	for v := range messages2 {
		fmt.Println(v)
	}
	// for {
	// 	select {
	// 	case res, ok := <-messages2:
	// 		if ok {

	// 			fmt.Println(res)
	// 		} else {
	// 			return
	// 		}
	// 	}
	// }

	// msg2 := <-messages2
	// fmt.Println(msg2)
	// msg3 := <-messages2
	// fmt.Println(msg3)
	// msg2 = <-messages2
	// fmt.Println(msg2)

}
