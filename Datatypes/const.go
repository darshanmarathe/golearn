package main

import "fmt"

//normal example
const (
	firstname  = "Darshan"
	middleName = "Narayan"
	lastName   = "Marathe"
)

const(
	age =20
	gender = "male"
)

//iota example
const (
	first  = 1 << (10 * iota)
	secend
	third
)

func main() {
	println("Normal example")
	fmt.Println(firstname)
	fmt.Println(middleName)
	fmt.Println(lastName)

	println("iota example")
	fmt.Println(first)
	fmt.Println(secend)
	fmt.Println(third)

	println("other exmaple")
	fmt.Println(age);
	fmt.Println(gender)
}
