package main

import (
	"context"
	"fmt"
	"time"
)

func doSomethingCool(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			fmt.Println("timed out")
			err := ctx.Err()
			fmt.Println("ERR: ", err)
			return
		default:
			fmt.Println("doing something cool")
		}

		time.Sleep(500 * time.Millisecond)
	}

}

func main() {
	fmt.Println("Go Context Tutorial")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	go doSomethingCool(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("oh no, I've exceeded the deadline")
	}

	time.Sleep(5 * time.Second)
}
