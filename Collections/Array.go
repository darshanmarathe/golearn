package main

import (
	"fmt"
)

type person struct {
	FirstName string
	LastName  string
}

func changeLocal(num *[5]int) {
	num[0] = 55
	fmt.Println("inside function ", num)

}

func main() {
	var arr [3]string
	arr[0] = "b1"
	arr[1] = "b2"
	arr[2] = "b2"

	marks := [4]int{10, 20, 30}

	fmt.Println("Demo for arrays", arr, len(arr))
	fmt.Println("Demo for arrays", marks, len(marks))

	a := [...]string{"USA", "China", "India", "Germany", "France"}
	b := a // a copy of a is assigned to b
	b[0] = "Singapore"
	fmt.Println("a is ", a)
	fmt.Println("b is ", b)

	darshan := &person{FirstName: "Darshan",
		LastName: "Marathe"}
	newDarshan := darshan
	newDarshan.FirstName = "New First Name"

	fmt.Println(darshan)
	fmt.Println(newDarshan)

	num := [...]int{5, 6, 7, 8, 8}
	fmt.Println("before passing to function ", num)
	changeLocal(&num) //num is passed by value
	fmt.Println("after passing to function ", num)

	for i := 0; i < len(num); i++ {
		fmt.Println(num[i])
	}
}
