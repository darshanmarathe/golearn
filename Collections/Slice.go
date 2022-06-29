package main

import (
	"fmt"
)

func main() {
	SliceDemo()
	return
	//Basic Slice creation
	var s []int
	s = make([]int, 3, 10)
	s[0] = 1
	s[1] = 10
	s[2] = 100

	//Short hand syntax
	s2 := []int{1, 10, 100, 1000}

	//Slicing the slice
	slice := s2[1:]

	//appending the slice
	slice = append(slice, 10000)

	//Deleting a slice
	nSlice := append(s2[1:], slice[:2]...)

	fmt.Println(s)
	fmt.Println(s2)
	fmt.Println(slice)
	fmt.Println(nSlice)
	fmt.Println("Demo for slices")
}

func SliceDemo() {
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4] //creates a slice from a[1] to a[3]
	b[2] = 22
	fmt.Println(a)
	fmt.Println(b)
	fibo := []int{0, 1, 1, 2, 3}

	fibo = append(fibo, 5)
	//Not possible with arrays
	//	a = append(a, 81)

	fmt.Println(fibo)

	i := make([]int, 5)
	i[6] = 100
	fmt.Println(i)
}
