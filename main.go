package main

import ( 
	"fmt"
)
const (
	message = "Hello go!"
)

var (
	item string
)

func main(){
	fmt.Println(message)
	fmt.Println(item)
}


func init(){
	item = "hello world!"
}