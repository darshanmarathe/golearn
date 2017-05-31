package main

import "fmt"

func main() {
	message := "hello word"
	var pointer *string = &message
	fmt.Println(message, pointer)
	fmt.Println(message, *pointer)

	*pointer = "New Code"
	fmt.Println(message)
}
