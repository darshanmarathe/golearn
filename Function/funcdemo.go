package main

import (
	"fmt"
)

func main()  {
	abuse := "some good message"
	Log(&abuse)
	Logs("Lets","develope","some","good","technologies","with","learning","go")
	println(abuse)
}

func Log(message *string){
	fmt.Println(*message)
}

func Logs(messages ...string){
	for _, msg := range messages {
		println(msg)
	}
}