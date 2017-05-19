package main

import "fmt"

//normal example
const (
	firstname ="Darshan"
	middleName = "Narayan"
	lastName = "Marathe"
)

//iota example
const(
	first = 1 << (10 * iota)
	secend
	third
)

func main()  {
	println("Normal example")
	fmt.Println(firstname)
	fmt.Println(middleName)
	fmt.Println(lastName)

	println("iota example")
	fmt.Println(first)
	fmt.Println(secend)
	fmt.Println(third)
}