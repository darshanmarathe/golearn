package main

import (
	"fmt"
)

func main() {
	darshan := Employee{}
	darshan.FirstName = "Darshan"
	darshan.LastName = "Marathe"
	darshan.fullname("hello")

}

type Employee struct {
	FirstName string
	LastName  string
}

func (emp *Employee) fullname(greet string) {
	fmt.Println(greet + " " + emp.FirstName + " " + emp.LastName)
}
