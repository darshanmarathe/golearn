package main

import (
	"fmt"
)

type Driver struct {
	FirstName string
	LastName  string
}

func main() {
	d := Driver{}
	d.FirstName = "Darshan"
	d.LastName = "Marathe"

	println(d.FirstName + " " + d.LastName)

	fmt.Println(d)

	a := Driver{"Ankit", "Goel"}
	println(a.FirstName + " " + a.LastName)

	fmt.Println(a)

	r := new(Driver)
	r.FirstName = "Raakesh"
	r.LastName = "Bhatia"
	println(r.FirstName + " " + r.LastName)

	fmt.Println(r)

}
