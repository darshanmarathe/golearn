package main

import (
	"fmt"
)

func main() {
	name := "Darshan"
	fmt.Println(SwitchDemo(name), " ", name)
	fmt.Println(SwitchDemo2(name), " ", name)
	fmt.Println(SwitchDemo3(name), " ", name)
	SwitchCheckType("Darshan")
	SwitchCheckType(20)
	SwitchCheckType(false)
	SwitchCheckType(1.220)
}

func SwitchDemo(name string) (title string) {

	switch name {
	case "Darshan":
		title = "Mr"
	case "Aditi":
		title = "Mrs"
	case "Shivansh":
		title = "Master"
	case "Dhruvi":
		title = "Miss"
	case "Nikhil":
		title = "Dr"
	default:
		title = "Dude"
	}
	return

}

//Fallthrogh demo 
func SwitchDemo2(name string) (title string) {

	switch name {
	case "Darshan":
		title = "Mr"
		fallthrough
	case "Aditi":
		title = "Mrs"
	case "Shivansh":
		title = "Master"
	case "Dhruvi":
		title = "Miss"
	case "Nikhil":
		title = "Dr"
	default:
		title = "Dude"
	}
	return
}

//Inline condition checking
func SwitchDemo3(name string) (title string) {

	switch {
	case name == "Darshan":
		title = "Mr"
	case name == "Aditi":
		title = "Mrs"
	case name == "Shivansh":
		title = "Master"
	case name == "Dhruvi":
		title = "Miss"
	case name == "Nikhil":
		title = "Dr"
	default:
		title = "Dude"
	}
	return
}

//Type check
func SwitchCheckType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	default:
		fmt.Println("unknown")

	}
}
