package main

import (
	"fmt"
)

func recoverFullName() {
	// if r := recover(); r != nil {
	// 	fmt.Println("recovered from ", r)
	// }
	r := recover()
	if r != nil {
		fmt.Println("recovered from ", r)
	}
}

func fullName(firstName *string, lastName *string) {
	defer recoverFullName()
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func slicePanic() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("recovered from ", r)
		}
	}()
	n := []int{5, 7, 4}
	fmt.Println(n[6])
	fmt.Println("normally returned from a")
}

func main() {
	defer fmt.Println("deferred call in main")
	firstName := "Elon"
	fullName(&firstName, nil)
	slicePanic()
	fmt.Println("returned normally from main")
}
