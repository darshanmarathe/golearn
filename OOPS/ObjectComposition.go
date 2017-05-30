package main

import (
	"fmt"
)

func main() {
	darshan := newOnlieEmployee("Darhan", "Marathe", "Narayan", 3000)

	darshan.fullname("hello")
	fmt.Println(darshan)
}

type Employee struct {
	FirstName string
	LastName  string
}

func (emp *Employee) fullname(greet string) {
	fmt.Println(greet + " " + emp.FirstName + " " + emp.LastName)
}

type OnlineEmployee struct {
	Employee
	MiddleName string
	salary     int
}

func newOnlieEmployee (fname string, lname string, mName string, sal int) *OnlineEmployee {
	oute := OnlineEmployee{}
	oute.FirstName = fname
	oute.LastName = lname
	oute.MiddleName = mName
	oute.salary = sal
	return &oute
}
